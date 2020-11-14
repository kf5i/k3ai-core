package commands

import (
	"fmt"
	"github.com/kf5i/k3ai-core/internal/k8s/kctl"
	"github.com/kf5i/k3ai-core/internal/plugins"
	"github.com/kf5i/k3ai-core/internal/shared"
)

const (
	//ApplyOperation operation apply
	ApplyOperation = "apply"
	//DeleteOperation operation delete
	DeleteOperation = "delete"
)

//HandlePlugin generic function to handle the plugins
func HandlePlugin(config kctl.Config, repo string, pluginName string, operation string) error {
	var plugin plugins.Plugin
	err := plugin.Encode(
		shared.NormalizeURL(repo, plugins.PluginDir, pluginName) + plugins.DefaultPluginFileName)
	if err != nil {
		return err
	}
	fmt.Printf("Plugin YAML content: %s, name: %s \n", plugin.Yaml, plugin.PluginName)
	switch operation {
	case ApplyOperation:
		{
			err = kctl.Apply(config, plugin, &kctl.CliWait{})
		}
	case DeleteOperation:
		err = kctl.Delete(config, plugin)
	}
	if err != nil {
		return err
	}
	return nil
}

//HandleGroup generic function to handle the groups
func HandleGroup(config kctl.Config, repo string, groupName string, operation string) error {
	var group plugins.Group
	err := group.Encode(
		shared.NormalizeURL(repo, plugins.GroupsDir, groupName) + plugins.DefaultGroupFileName)
	if err != nil {
		return err
	}
	fmt.Printf("Plugin YAML group-name: %s \n", group.GroupName)

	for _, plugin := range group.Plugins {
		err := HandlePlugin(config, repo, plugin.Name, operation)
		if err != nil {
			fmt.Printf("Error during handling plugin, error: %s", err)
		}
	}

	for _, inlinePlugin := range group.InlinePlugins {
		switch operation {
		case ApplyOperation:
			{
				err = kctl.Apply(config, inlinePlugin, &kctl.CliWait{})
				if err != nil {
					fmt.Printf("Error during handling Apply inline Plugin, error: %s", err)
				}
			}

		case DeleteOperation:
			err = kctl.Delete(config, inlinePlugin)
			if err != nil {
				fmt.Printf("Error during handling Delete inline Plugin, error: %s", err)
			}
		}
	}

	return nil
}
