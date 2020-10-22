package main

import (
	"fmt"
	"github.com/kf5i/k3ai-core/pkg/k8s/kctl"
	"github.com/kf5i/k3ai-core/pkg/plugins"
)

func main() {
	pluginList, _ := plugins.GetPluginList()
	fmt.Printf("Plugin list: %s\n", pluginList)

	githubContents, _ := plugins.GetPluginYamlS("argo")
	for _, githubContent := range *githubContents {
		fmt.Printf("Files: %s, Path %s \n", githubContent.Name, githubContent.Path)
		var pluginSpec plugins.PluginSpec
		pluginSpec.Encode(githubContent.Path)
		fmt.Printf("Plugin YAML content: %s, name: %s \n", pluginSpec.Files, pluginSpec.PluginName)
		fmt.Println("Going to Apply the Apply")
		kctl.ApplyFiles(pluginSpec, nil)
	}

}
