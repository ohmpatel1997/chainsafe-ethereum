package dagpb_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-blockservice"
	"github.com/ipfs/go-datastore"
	dss "github.com/ipfs/go-datastore/sync"
	bstore "github.com/ipfs/go-ipfs-blockstore"
	chunker "github.com/ipfs/go-ipfs-chunker"
	offline "github.com/ipfs/go-ipfs-exchange-offline"
	files "github.com/ipfs/go-ipfs-files"
	ipldformat "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-merkledag"
	unixfile "github.com/ipfs/go-unixfs/file"
	"github.com/ipfs/go-unixfs/importer/balanced"
	ihelper "github.com/ipfs/go-unixfs/importer/helpers"
	ipld "github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/ipld/go-ipld-prime/traversal"
	"github.com/ipld/go-ipld-prime/traversal/selector"
	"github.com/ipld/go-ipld-prime/traversal/selector/builder"
	. "github.com/warpfork/go-wish"

	dagpb "github.com/ipld/go-ipld-prime-proto"
)

const UnixfsChunkSize uint64 = 1 << 10
const UnixfsLinksPerLevel = 1024

func makeLoader(bs bstore.Blockstore) ipld.Loader {
	return func(lnk ipld.Link, lnkCtx ipld.LinkContext) (io.Reader, error) {
		c, ok := lnk.(cidlink.Link)
		if !ok {
			return nil, errors.New("Incorrect Link Type")
		}
		// read block from one store
		block, err := bs.Get(c.Cid)
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(block.RawData()), nil
	}
}

func makeStorer(bs bstore.Blockstore) ipld.Storer {
	return func(lnkCtx ipld.LinkContext) (io.Writer, ipld.StoreCommitter, error) {
		var buf bytes.Buffer
		var committer ipld.StoreCommitter = func(lnk ipld.Link) error {
			c, ok := lnk.(cidlink.Link)
			if !ok {
				return errors.New("Incorrect Link Type")
			}
			block, err := blocks.NewBlockWithCid(buf.Bytes(), c.Cid)
			if err != nil {
				return err
			}
			return bs.Put(block)
		}
		return &buf, committer, nil
	}
}

// What this test does:
// - Construct a blockstore + dag service
// - Import a file to UnixFS v1
// - Reload the UnixFS v1 DAG in go-ipld-prime
// - Traverse the dag with a selector
// - Re-encode any protobuf or raw nodes encountered along
// the way to a new block store
// - Load the file from the new block store using the
// existing UnixFS v1 file reader
// - Verify the bytes match the original

func TestUnixFSSelectorCopy(t *testing.T) {
	ctx := context.Background()

	// make a blockstore and dag service
	bs1 := bstore.NewBlockstore(dss.MutexWrap(datastore.NewMapDatastore()))
	dagService1 := merkledag.NewDAGService(blockservice.New(bs1, offline.Exchange(bs1)))

	// make a second blockstore
	bs2 := bstore.NewBlockstore(dss.MutexWrap(datastore.NewMapDatastore()))

	// read in a fixture file
	path, err := filepath.Abs(filepath.Join("fixtures", "lorem.txt"))
	Wish(t, err, ShouldEqual, nil)

	f, err := os.Open(path)
	Wish(t, err, ShouldEqual, nil)

	var buf bytes.Buffer
	tr := io.TeeReader(f, &buf)
	file := files.NewReaderFile(tr)

	// import to UnixFS
	bufferedDS := ipldformat.NewBufferedDAG(ctx, dagService1)

	params := ihelper.DagBuilderParams{
		Maxlinks:   UnixfsLinksPerLevel,
		RawLeaves:  true,
		CidBuilder: nil,
		Dagserv:    bufferedDS,
	}

	db, err := params.New(chunker.NewSizeSplitter(file, int64(UnixfsChunkSize)))
	Wish(t, err, ShouldEqual, nil)
	nd, err := balanced.Layout(db)
	Wish(t, err, ShouldEqual, nil)
	err = bufferedDS.Commit()
	Wish(t, err, ShouldEqual, nil)

	// save the original files bytes
	origBytes := buf.Bytes()

	// setup an IPLD loader for blockstore 1
	loader := makeLoader(bs1)

	// setup an IPLD storer for blockstore 2
	storer := makeStorer(bs2)

	// load the root of the UnixFS dag in go-ipld-prime
	clink := cidlink.Link{Cid: nd.Cid()}
	nb := dagpb.Type.PBNode.NewBuilder()
	err = clink.Load(ctx, ipld.LinkContext{}, nb, loader)
	Wish(t, err, ShouldEqual, nil)

	primeNd := nb.Build()
	// get a protobuf link builder
	pbLinkBuilder := clink.LinkBuilder()

	// get a raw link builder
	links, err := primeNd.LookupByString("Links")
	Wish(t, err, ShouldEqual, nil)
	link, err := links.LookupByIndex(0)
	Wish(t, err, ShouldEqual, nil)
	rawLink, err := link.LookupByString("Hash")
	Wish(t, err, ShouldEqual, nil)
	rawLinkLink, err := rawLink.AsLink()
	Wish(t, err, ShouldEqual, nil)
	rawLinkBuilder := rawLinkLink.LinkBuilder()

	// setup a node builder chooser with DagPB + Raw support
	var defaultChooser traversal.LinkTargetNodePrototypeChooser = dagpb.AddDagPBSupportToChooser(func(ipld.Link, ipld.LinkContext) (ipld.NodePrototype, error) {
		return basicnode.Prototype.Any, nil
	})

	// create a selector for the whole UnixFS dag
	ssb := builder.NewSelectorSpecBuilder(basicnode.Prototype.Any)

	allSelector, err := ssb.ExploreRecursive(selector.RecursionLimitNone(),
		ssb.ExploreAll(ssb.ExploreRecursiveEdge())).Selector()
	Wish(t, err, ShouldEqual, nil)

	// execute the traversal
	err = traversal.Progress{
		Cfg: &traversal.Config{
			LinkLoader:                     loader,
			LinkTargetNodePrototypeChooser: defaultChooser,
		},
	}.WalkAdv(primeNd, allSelector, func(pg traversal.Progress, nd ipld.Node, r traversal.VisitReason) error {
		// for each node encountered, check if it's a DabPB Node or a Raw Node and if so
		// encode and store it in the new blockstore
		pbNode, ok := nd.(dagpb.PBNode)
		if ok {
			_, err := pbLinkBuilder.Build(ctx, ipld.LinkContext{}, pbNode, storer)
			return err
		}
		rawNode, ok := nd.(dagpb.RawNode)
		if ok {
			_, err := rawLinkBuilder.Build(ctx, ipld.LinkContext{}, rawNode, storer)
			return err
		}
		return nil
	})

	// verify traversal was successful
	Wish(t, err, ShouldEqual, nil)

	// setup a DagService for the second block store
	dagService2 := merkledag.NewDAGService(blockservice.New(bs2, offline.Exchange(bs2)))

	// load the root of the UnixFS DAG from the new blockstore
	otherNode, err := dagService2.Get(ctx, nd.Cid())
	Wish(t, err, ShouldEqual, nil)

	// Setup a UnixFS file reader
	n, err := unixfile.NewUnixfsFile(ctx, dagService2, otherNode)
	Wish(t, err, ShouldEqual, nil)
	fn, ok := n.(files.File)
	Wish(t, ok, ShouldEqual, true)

	// Read the bytes for the UnixFS File
	finalBytes, err := ioutil.ReadAll(fn)
	Wish(t, err, ShouldEqual, nil)

	// verify original bytes match final bytes!
	Wish(t, origBytes, ShouldEqual, finalBytes)

}
