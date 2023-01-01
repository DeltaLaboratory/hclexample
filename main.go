package main

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclsimple"
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
	fmt.Printf("result : %#v", stru)
}
