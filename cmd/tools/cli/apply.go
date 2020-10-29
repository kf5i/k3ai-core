package cli

import (
	"fmt"

	"github.com/kf5i/k3ai-core/internal/k8s/kctl"
	"github.com/kf5i/k3ai-core/internal/plugins"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply <plugin_name>",
	Short: "Apply the plugin",
	//	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		config := newConfig()
		plugin, _ := cmd.Flags().GetString("plugin")
		if plugin != "" {
			return singlePlugin(config, plugin)
		}
		group, _ := cmd.Flags().GetString("group")
		if group != "" {
			return pluginsGroup(config, group)

		}

		return nil
	},
}

func pluginsGroup(config kctl.Config, destination string) error {
	pluginsGroupSpec, err := plugins.LoadPluginsGroupSpecFormFile(pluginsGroupRepoURI + "/plugins_group.yaml")
	if err != nil {
		return err
	}

	for _, pluginGroupSpec := range pluginsGroupSpec.Plugins {
		if pluginGroupSpec.Enabled == true {
			err := singlePlugin(config, pluginGroupSpec.Name)
			if err != nil {
				return err
			}
		}
	}

	return nil

}

func singlePlugin(config kctl.Config, destination string) error {
	pluginSpecList, err := plugins.PluginYamls(pluginRepoURI, destination)
	if err != nil {
		return err
	}
	for _, pluginSpec := range pluginSpecList {
		fmt.Printf("Plugin YAML content: %s, name: %s \n", pluginSpec.Yaml, pluginSpec.PluginName)
		err = kctl.Apply(config, pluginSpec, nil)
		if err != nil {
			return err
		}
	}
	return nil
}
