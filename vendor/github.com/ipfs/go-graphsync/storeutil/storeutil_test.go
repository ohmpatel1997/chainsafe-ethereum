package storeutil

import (
	"io/ioutil"
	"testing"

	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-datastore"
	dss "github.com/ipfs/go-datastore/sync"
	bstore "github.com/ipfs/go-ipfs-blockstore"
	ipld "github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/stretchr/testify/require"

	"github.com/ipfs/go-graphsync/testutil"
)

func TestLoader(t *testing.T) {
	store := bstore.NewBlockstore(dss.MutexWrap(datastore.NewMapDatastore()))
	blk := testutil.GenerateBlocksOfSize(1, 1000)[0]
	err := store.Put(blk)
	require.NoError(t, err, "Unable to put block to store")
	loader := LoaderForBlockstore(store)
	data, err := loader(cidlink.Link{Cid: blk.Cid()}, ipld.LinkContext{})
	require.NoError(t, err, "Unable to load block with loader")
	bytes, err := ioutil.ReadAll(data)
	require.NoError(t, err, "Unable to read bytes from reader returned by loader")
	_, err = blocks.NewBlockWithCid(bytes, blk.Cid())
	require.NoError(t, err, "Did not return correct block with loader")
}

func TestStorer(t *testing.T) {
	store := bstore.NewBlockstore(dss.MutexWrap(datastore.NewMapDatastore()))
	blk := testutil.GenerateBlocksOfSize(1, 1000)[0]
	storer := StorerForBlockstore(store)
	buffer, commit, err := storer(ipld.LinkContext{})
	require.NoError(t, err, "Unable to setup buffer")
	_, err = buffer.Write(blk.RawData())
	require.NoError(t, err, "Unable to write data to buffer")
	err = commit(cidlink.Link{Cid: blk.Cid()})
	require.NoError(t, err, "Unable to commit with storer function")
	_, err = store.Get(blk.Cid())
	require.NoError(t, err, "Block not written to store")
}
