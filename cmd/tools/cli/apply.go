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
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		config := newConfig()
		plugin, _ := cmd.Flags().GetString(plugins.PluginType)
		if plugin != "" {
			return applyPlugin(config, plugin)
		}
		group, _ := cmd.Flags().GetString(plugins.GroupType)
		if group != "" {
			return applyGroup(config, group)

		}

		return nil
	},
}

func applyGroup(config kctl.Config, groupName string) error {
	var groups plugins.Groups
	pluginsGroupSpec, err := groups.Encode(pluginsGroupRepoURI, groupName)
	if err != nil {
		return err
	}

	for _, group := range pluginsGroupSpec.Groups {
		for _, plugin := range group.Plugins {
			if plugin.Enabled == true {
				err := applyPlugin(config, plugin.Name)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil

}

func applyPlugin(config kctl.Config, pluginName string) error {
	var plugins plugins.Plugins
	pluginSpecList, err := plugins.Encode(pluginRepoURI, pluginName)
	if err != nil {
		return err
	}
	for _, pluginSpec := range pluginSpecList.Plugins {
		fmt.Printf("Plugin YAML content: %s, name: %s \n", pluginSpec.Yaml, pluginSpec.PluginName)
		err = kctl.Apply(config, pluginSpec, nil)
		if err != nil {
			return err
		}
	}
	return nil
}
