package cli

import (
	"github.com/kf5i/k3ai-core/internal/k8s/kctl"
	"github.com/kf5i/k3ai-core/internal/plugins"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <plugin_name>",
	Short: "Delete the plugin",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		config := newConfig()
		pluginName := args[0]
		pluginSpecList, err := plugins.GetPluginYamls(pluginRepoUri, pluginName)
		if err != nil {
			return err
		}
		for _, pluginSpec := range pluginSpecList {
			err = kctl.Delete(config, pluginSpec)
			if err != nil {
				return err
			}
		}
		return nil
	},
}
