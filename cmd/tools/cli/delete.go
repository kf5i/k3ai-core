package cli

import (
	"fmt"
	"github.com/kf5i/k3ai-core/internal/k8s/kctl"
	"github.com/kf5i/k3ai-core/internal/plugins"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <plugin_name>",
	Short: "Delete the plugin",
	//Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		config := newConfig()

		plugin, _ := cmd.Flags().GetString(plugins.PluginType)
		if plugin != "" {
			return deletePlugin(config, plugin)
		}
		group, _ := cmd.Flags().GetString(plugins.GroupType)
		if group != "" {
			return deleteGroup(config, group)

		}
		return nil
	},
}

func deleteGroup(config kctl.Config, groupName string) error {
	var g plugins.Group
	group, err := g.Encode(plugins.NormalizePath(pluginsGroupRepoURI,
		groupName,
		plugins.DefaultGroupFileName))
	if err != nil {
		return err
	}

	for _, pluginItem := range group.Plugins {
		if pluginItem.Enabled == true {
			err := deletePlugin(config, pluginItem.Name)
			if err != nil {
				return err
			}
		}
	}

	return nil

}

func deletePlugin(config kctl.Config, pluginName string) error {
	var plugins plugins.Plugins
	pluginSpecList, err := plugins.Encode(pluginRepoURI, pluginName)
	if err != nil {
		return err
	}
	for _, pluginItem := range pluginSpecList.Plugins {
		fmt.Printf("Plugin YAML content: %s, name: %s \n", pluginItem.Yaml, pluginItem.PluginName)
		err = kctl.Delete(config, pluginItem)
		if err != nil {
			return err
		}
	}
	return nil
}
