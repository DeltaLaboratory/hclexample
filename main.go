package main

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

type SomeStructure struct {
	SomeValue  string                `hcl:"some_value"`
	SomeStruct SomeEmbeddedStructure `hcl:"some_struct,block"`
}

type SomeEmbeddedStructure struct {
	SomeLabel string   `hcl:"some_label,label"`
	SomeValue string   `hcl:"some_value"`
	SomeArray []string `hcl:"some_array"`
}

func main() {
	var stru SomeStructure
	if err := hclsimple.DecodeFile("example.hcl", nil, &stru); err != nil {
		panic(err)
	}
	fmt.Printf("decode result : %#v", stru)
	f := hclwrite.NewEmptyFile()
	b := f.Body()
	b.SetAttributeValue("some_value", cty.StringVal("value"))
	b.AppendNewline()
	ebb := b.AppendNewBlock("some_struct", []string{"some_label"})
	eb := ebb.Body()
	eb.SetAttributeValue("some_value", cty.StringVal("value!"))
	eb.SetAttributeValue("some_array", cty.ListVal([]cty.Value{cty.StringVal("element1"), cty.StringVal("element2")}))
	fmt.Printf("%s\n", f.Bytes())
}
