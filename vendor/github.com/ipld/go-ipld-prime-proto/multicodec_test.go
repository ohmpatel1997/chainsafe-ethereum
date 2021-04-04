package dagpb_test

import (
	"bytes"
	"testing"

	dag "github.com/ipfs/go-merkledag"
	dagpb "github.com/ipld/go-ipld-prime-proto"
	. "github.com/warpfork/go-wish"
)

func TestRoundTripRaw(t *testing.T) {
	randBytes := randomBytes(256)
	rawNode, err := makeRawNode(randBytes)
	Wish(t, err, ShouldEqual, nil)
	t.Run("encoding", func(t *testing.T) {
		var buf bytes.Buffer
		err := dagpb.RawEncoder(rawNode, &buf)
		Wish(t, err, ShouldEqual, nil)
		Wish(t, buf.Bytes(), ShouldEqual, randBytes)
	})
	t.Run("decoding", func(t *testing.T) {
		buf := bytes.NewBuffer(randBytes)
		nb := dagpb.Type.RawNode.NewBuilder()
		err := dagpb.RawDecoder(nb, buf)
		Wish(t, err, ShouldEqual, nil)
		rawNode2 := nb.Build()
		Wish(t, rawNode2, ShouldEqual, rawNode)
	})
}

func TestRoundTripProtbuf(t *testing.T) {
	a := dag.NewRawNode([]byte("aaaa"))
	b := dag.NewRawNode([]byte("bbbb"))
	c := dag.NewRawNode([]byte("cccc"))

	nd1 := &dag.ProtoNode{}
	nd1.AddNodeLink("cat", a)

	nd2 := &dag.ProtoNode{}
	nd2.AddNodeLink("first", nd1)
	nd2.AddNodeLink("dog", b)

	nd3 := &dag.ProtoNode{}
	nd3.AddNodeLink("second", nd2)
	nd3.AddNodeLink("bear", c)

	data := nd3.RawData()
	ibuf := bytes.NewBuffer(data)
	nb := dagpb.Type.PBNode.NewBuilder()
	err := dagpb.PBDecoder(nb, ibuf)
	Wish(t, err, ShouldEqual, nil)
	pbNode := nb.Build()
	t.Run("encode/decode equivalency", func(t *testing.T) {
		var buf bytes.Buffer
		err := dagpb.PBEncoder(pbNode, &buf)
		Wish(t, err, ShouldEqual, nil)
		nb = dagpb.Type.PBNode.NewBuilder()
		err = dagpb.PBDecoder(nb, &buf)
		pbNode2 := nb.Build()
		Wish(t, err, ShouldEqual, nil)
		Wish(t, pbNode2, ShouldEqual, pbNode)
	})
}
