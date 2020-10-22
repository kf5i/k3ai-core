package main

import (
	"fmt"
	"github.com/kf5i/k3ai-core/pkg/k8s/kctl"
	"github.com/kf5i/k3ai-core/pkg/plugins"
)

type TestType struct {
	testField int
}

func main() {
	var encode plugins.PluginSpec
	encode.Encode()
	fmt.Printf("%s\n", encode.Files)

	s := []string{"aaa"}

	kctl.ApplyFiles(s, nil)

	p, _ := plugins.GetPluginList()

	fmt.Printf("List %s\n", p)

	pf, _ := plugins.GetPluginYamlS("argo")

	fmt.Printf("List %s\n", pf)

	//kctl.ApplyFiles(encode.Files)
}
