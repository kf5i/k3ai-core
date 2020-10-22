package main

import (
	"fmt"

	"github.com/kf5i/k3ai-core/pkg/k8s/kctl"
	"github.com/kf5i/k3ai-core/pkg/plugins"
)

func main() {
	pluginList, _ := plugins.GetPluginList()
	fmt.Printf("Plugin list: %s\n", pluginList)

	pluginSpecList, _ := plugins.GetPluginYamls("argo")
	for _, pluginSpec := range *pluginSpecList {
		//fmt.Printf("Files: %s, Path %s \n", githubContent.Name, githubContent.Path)
		fmt.Printf("Plugin YAML content: %s, name: %s \n", pluginSpec.Files, pluginSpec.PluginName)
		fmt.Println("Going to Apply the Apply")
		kctl.ApplyFiles(pluginSpec, nil)
	}

}
