package main

import (
	"os/exec"

	"github.com/ipld/go-ipld-prime/schema"
	gengo "github.com/ipld/go-ipld-prime/schema/gen/go"
)

func main() {

	ts := schema.TypeSystem{}
	ts.Init()
	adjCfg := &gengo.AdjunctCfg{}

	pkgName := "dagpb"

	ts.Accumulate(schema.SpawnString("String"))
	ts.Accumulate(schema.SpawnInt("Int"))
	ts.Accumulate(schema.SpawnLink("Link"))
	ts.Accumulate(schema.SpawnBytes("Bytes"))

	ts.Accumulate(schema.SpawnStruct("PBLink",
		[]schema.StructField{
			schema.SpawnStructField("Hash", "Link", true, false),
			schema.SpawnStructField("Name", "String", true, false),
			schema.SpawnStructField("Tsize", "Int", true, false),
		},
		schema.SpawnStructRepresentationMap(nil),
	))
	ts.Accumulate(schema.SpawnList("PBLinks", "PBLink", false))
	ts.Accumulate(schema.SpawnStruct("PBNode",
		[]schema.StructField{
			schema.SpawnStructField("Links", "PBLinks", false, false),
			schema.SpawnStructField("Data", "Bytes", false, false),
		},
		schema.SpawnStructRepresentationMap(nil),
	))
	ts.Accumulate(schema.SpawnBytes("RawNode"))
	gengo.Generate(".", pkgName, ts, adjCfg)
	exec.Command("go", "fmt").Run()
}
