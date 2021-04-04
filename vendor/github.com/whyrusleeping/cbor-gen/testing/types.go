package testing

import (
	cbg "github.com/whyrusleeping/cbor-gen"
)

const Thingc = 3

type NaturalNumber uint64

type SignedArray struct {
	Signed []uint64
}

type SimpleTypeOne struct {
	Foo    string
	Value  uint64
	Binary []byte
	Signed int64
}

type SimpleTypeTwo struct {
	Stuff        *SimpleTypeTwo
	Others       []uint64
	SignedOthers []int64
	Test         [][]byte
	Dog          string
	Numbers      []NaturalNumber
	Pizza        *uint64
	PointyPizza  *NaturalNumber
	Arrrrrghay   [Thingc]SimpleTypeOne
}

type SimpleTypeTree struct {
	Stuff                            *SimpleTypeTree
	Stufff                           *SimpleTypeTwo
	Others                           []uint64
	Test                             [][]byte
	Dog                              string
	SixtyThreeBitIntegerWithASignBit int64
	NotPizza                         *uint64
}

type DeferredContainer struct {
	Stuff    *SimpleTypeOne
	Deferred *cbg.Deferred
	Value    uint64
}
