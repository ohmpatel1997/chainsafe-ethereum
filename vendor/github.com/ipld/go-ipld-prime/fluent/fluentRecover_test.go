package fluent_test

import (
	"testing"

	. "github.com/warpfork/go-wish"

	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/fluent"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
)

func TestRecover(t *testing.T) {
	t.Run("simple build error should capture", func(t *testing.T) {
		Wish(t,
			fluent.Recover(func() {
				fluent.MustBuild(basicnode.Prototype__String{}, func(fna fluent.NodeAssembler) {
					fna.AssignInt(9)
				})
				t.Fatal("should not be reached")
			}),
			ShouldEqual,
			fluent.Error{ipld.ErrWrongKind{TypeName: "string", MethodName: "AssignInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_String}},
		)
	})
	t.Run("correct build should return nil", func(t *testing.T) {
		Wish(t,
			fluent.Recover(func() {
				fluent.MustBuild(basicnode.Prototype__String{}, func(fna fluent.NodeAssembler) {
					fna.AssignString("fine")
				})
			}),
			ShouldEqual,
			nil,
		)
	})
	t.Run("other panics should continue to rise", func(t *testing.T) {
		Wish(t,
			func() (r interface{}) {
				defer func() { r = recover() }()
				fluent.Recover(func() {
					panic("fuqawds")
				})
				return
			}(),
			ShouldEqual,
			"fuqawds",
		)
	})
}
