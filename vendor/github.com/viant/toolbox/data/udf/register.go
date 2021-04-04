package udf

import "github.com/viant/toolbox/data"



//Register registers defined in this package UDFs
func Register(aMap data.Map) {
	aMap.Put("AsInt", AsInt)
	aMap.Put("AsString", AsString)
	aMap.Put("AsFloat", AsFloat)
	aMap.Put("AsFloat32", AsFloat32)
	aMap.Put("AsFloat32Ptr", AsFloat32Ptr)

	aMap.Put("AsBool", AsBool)
	aMap.Put("AsMap", AsMap)
	aMap.Put("AsData", AsData)
	aMap.Put("AsCollection", AsCollection)
	aMap.Put("AsJSON", AsJSON)
	aMap.Put("Type", Type)
	aMap.Put("Join", Join)
	aMap.Put("Split", Split)
	aMap.Put("Keys", Keys)
	aMap.Put("Values", Values)
	aMap.Put("Length", Length)
	aMap.Put("Len", Length)
	aMap.Put("IndexOf", IndexOf)
	aMap.Put("FormatTime", FormatTime)
	aMap.Put("QueryEscape", QueryEscape)
	aMap.Put("QueryUnescape", QueryUnescape)
	aMap.Put("Base64Encode", Base64Encode)
	aMap.Put("Base64Decode", Base64Decode)
	aMap.Put("Base64DecodeText", Base64DecodeText)
	aMap.Put("TrimSpace", TrimSpace)
	aMap.Put("Elapsed", Elapsed)
	aMap.Put("Sum", Sum)
	aMap.Put("Count", Count)
	aMap.Put("AsNumber", AsNumber)
	aMap.Put("Select", Select)
	aMap.Put("Rand", Rand)
	aMap.Put("Concat", Concat)
	aMap.Put("Merge", Merge)
	aMap.Put("AsStringMap", AsStringMap)
}
