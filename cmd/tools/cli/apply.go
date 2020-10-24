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
		pluginName := args[0]
		pluginList, err := plugins.GetPluginList()
		if err != nil {
			return err
		}
		fmt.Printf("Plugin list: %s\n", pluginList)

		pluginSpecList, _ := plugins.GetPluginYamls(pluginName)
		for _, pluginSpec := range *pluginSpecList {
			fmt.Printf("Plugin YAML content: %s, name: %s \n", pluginSpec.Yaml, pluginSpec.PluginName)
			kctl.Apply(config, pluginSpec, nil)
		}
		return nil
	},
}
