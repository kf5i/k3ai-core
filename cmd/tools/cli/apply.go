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
		pluginSpecList, err := plugins.GetPluginYamls(pluginRepoUri, pluginName)
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
	},
}
