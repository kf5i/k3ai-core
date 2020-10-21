package main

import (
	"fmt"
	"github.com/kf5i/k3ai-core/pkg/plugins"
)

func main() {
	//var encode plugins.PluginSpec
	//encode.Encode()
	//fmt.Printf("%s\n", encode.Files)

	p, _ := plugins.GetPluginList()

	fmt.Printf("List %s\n", p)

	pf, _ := plugins.GetPluginYamls("argo")

	fmt.Printf("List %s\n", pf)


	//kctl.ApplyFiles(encode.Files)
}
