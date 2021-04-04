package mixins

import (
	"io"

	ipld "github.com/ipld/go-ipld-prime"
)

type IntTraits struct {
	PkgName    string
	TypeName   string // see doc in kindTraitsGenerator
	TypeSymbol string // see doc in kindTraitsGenerator
}

func (IntTraits) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Int
}
func (g IntTraits) EmitNodeMethodReprKind(w io.Writer) {
	doTemplate(`
		func ({{ .TypeSymbol }}) ReprKind() ipld.ReprKind {
			return ipld.ReprKind_Int
		}
	`, w, g)
}
func (g IntTraits) EmitNodeMethodLookupByString(w io.Writer) {
	kindTraitsGenerator{g.PkgName, g.TypeName, g.TypeSymbol, ipld.ReprKind_Int}.emitNodeMethodLookupByString(w)
}
func (g IntTraits) EmitNodeMethodLookupByNode(w io.Writer) {
	kindTraitsGenerator{g.PkgName, g.TypeName, g.TypeSymbol, ipld.ReprKind_Int}.emitNodeMethodLookupByNode(w)
}
func (g IntTraits) EmitNodeMethodLookupByIndex(w io.Writer) {
	kindTraitsGenerator{g.PkgName, g.TypeName, g.TypeSymbol, ipld.ReprKind_Int}.emitNodeMethodLookupByIndex(w)
}
func (g IntTraits) EmitNodeMethodLookupBySegment(w io.Writer) {
	kindTraitsGenerator{g.PkgName, g.TypeName, g.TypeSymbol, ipld.ReprKind_Int}.emitNodeMethodLookupBySegment(w)
}
func (g IntTraits) EmitNodeMethodMapIterator(w io.Writer) {
	kindTraitsGenerator{g.PkgName, g.TypeName, g.TypeSymbol, ipld.ReprKind_Int}.emitNodeMethodMapIterator(w)
}
func (g IntTraits) EmitNodeMethodListIterator(w io.Writer) {
	kindTraitsGenerator{g.PkgName, g.TypeName, g.TypeSymbol, ipld.ReprKind_Int}.emitNodeMethodListIterator(w)
}
func (g IntTraits) EmitNodeMethodLength(w io.Writer) {
	kindTraitsGenerator{g.PkgName, g.TypeName, g.TypeSymbol, ipld.ReprKind_Int}.emitNodeMethodLength(w)
}
func (g IntTraits) EmitNodeMethodIsAbsent(w io.Writer) {
	kindTraitsGenerator{g.PkgName, g.TypeName, g.TypeSymbol, ipld.ReprKind_Int}.emitNodeMethodIsAbsent(w)
}
func (g IntTraits) EmitNodeMethodIsNull(w io.Writer) {
	kindTraitsGenerator{g.PkgName, g.TypeName, g.TypeSymbol, ipld.ReprKind_Int}.emitNodeMethodIsNull(w)
}
func (g IntTraits) EmitNodeMethodAsBool(w io.Writer) {
	kindTraitsGenerator{g.PkgName, g.TypeName, g.TypeSymbol, ipld.ReprKind_Int}.emitNodeMethodAsBool(w)
}
func (g IntTraits) EmitNodeMethodAsFloat(w io.Writer) {
	kindTraitsGenerator{g.PkgName, g.TypeName, g.TypeSymbol, ipld.ReprKind_Int}.emitNodeMethodAsFloat(w)
}
func (g IntTraits) EmitNodeMethodAsString(w io.Writer) {
	kindTraitsGenerator{g.PkgName, g.TypeName, g.TypeSymbol, ipld.ReprKind_Int}.emitNodeMethodAsString(w)
}
func (g IntTraits) EmitNodeMethodAsBytes(w io.Writer) {
	kindTraitsGenerator{g.PkgName, g.TypeName, g.TypeSymbol, ipld.ReprKind_Int}.emitNodeMethodAsBytes(w)
}
func (g IntTraits) EmitNodeMethodAsLink(w io.Writer) {
	kindTraitsGenerator{g.PkgName, g.TypeName, g.TypeSymbol, ipld.ReprKind_Int}.emitNodeMethodAsLink(w)
}

type IntAssemblerTraits struct {
	PkgName       string
	TypeName      string // see doc in kindAssemblerTraitsGenerator
	AppliedPrefix string // see doc in kindAssemblerTraitsGenerator
}

func (IntAssemblerTraits) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Int
}
func (g IntAssemblerTraits) EmitNodeAssemblerMethodBeginMap(w io.Writer) {
	kindAssemblerTraitsGenerator{g.PkgName, g.TypeName, g.AppliedPrefix, ipld.ReprKind_Int}.emitNodeAssemblerMethodBeginMap(w)
}
func (g IntAssemblerTraits) EmitNodeAssemblerMethodBeginList(w io.Writer) {
	kindAssemblerTraitsGenerator{g.PkgName, g.TypeName, g.AppliedPrefix, ipld.ReprKind_Int}.emitNodeAssemblerMethodBeginList(w)
}
func (g IntAssemblerTraits) EmitNodeAssemblerMethodAssignNull(w io.Writer) {
	kindAssemblerTraitsGenerator{g.PkgName, g.TypeName, g.AppliedPrefix, ipld.ReprKind_Int}.emitNodeAssemblerMethodAssignNull(w)
}
func (g IntAssemblerTraits) EmitNodeAssemblerMethodAssignBool(w io.Writer) {
	kindAssemblerTraitsGenerator{g.PkgName, g.TypeName, g.AppliedPrefix, ipld.ReprKind_Int}.emitNodeAssemblerMethodAssignBool(w)
}
func (g IntAssemblerTraits) EmitNodeAssemblerMethodAssignFloat(w io.Writer) {
	kindAssemblerTraitsGenerator{g.PkgName, g.TypeName, g.AppliedPrefix, ipld.ReprKind_Int}.emitNodeAssemblerMethodAssignFloat(w)
}
func (g IntAssemblerTraits) EmitNodeAssemblerMethodAssignString(w io.Writer) {
	kindAssemblerTraitsGenerator{g.PkgName, g.TypeName, g.AppliedPrefix, ipld.ReprKind_Int}.emitNodeAssemblerMethodAssignString(w)
}
func (g IntAssemblerTraits) EmitNodeAssemblerMethodAssignBytes(w io.Writer) {
	kindAssemblerTraitsGenerator{g.PkgName, g.TypeName, g.AppliedPrefix, ipld.ReprKind_Int}.emitNodeAssemblerMethodAssignBytes(w)
}
func (g IntAssemblerTraits) EmitNodeAssemblerMethodAssignLink(w io.Writer) {
	kindAssemblerTraitsGenerator{g.PkgName, g.TypeName, g.AppliedPrefix, ipld.ReprKind_Int}.emitNodeAssemblerMethodAssignLink(w)
}
func (g IntAssemblerTraits) EmitNodeAssemblerMethodPrototype(w io.Writer) {
	kindAssemblerTraitsGenerator{g.PkgName, g.TypeName, g.AppliedPrefix, ipld.ReprKind_Int}.emitNodeAssemblerMethodPrototype(w)
}
