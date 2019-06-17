package main

import (
	"flag"

	"github.com/golang/glog"
	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protokit"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := protokit.RunPlugin(new(plugin)); err != nil {
		glog.Fatal(err)
	}
}

type plugin struct{}

func (p *plugin) Generate(req *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	descriptors := protokit.ParseCodeGenRequest(req)
	resp := new(plugin_go.CodeGeneratorResponse)

	BuildTypeNameMap(descriptors)
	Merge(resp, descriptors)
	return resp, nil
}
