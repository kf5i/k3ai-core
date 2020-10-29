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
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		config := newConfig()
		commandType := args[0]
		destination := args[1]

		switch commandType {
		case "-p":
			err2, done := singlePlugin(config, destination)
			if done {
				return err2
			}
		case "-g":
			err2, done := pluginsGroup(config, destination)
			if done {
				return err2
			}
		}

		return nil
	},
}

func pluginsGroup(config kctl.Config, destination string) (error, bool) {
	return nil, false

}

func singlePlugin(config kctl.Config, destination string) (error, bool) {
	pluginSpecList, err := plugins.PluginYamls(pluginRepoURI, destination)
	if err != nil {
		return err, true
	}
	for _, pluginSpec := range pluginSpecList {
		fmt.Printf("Plugin YAML content: %s, name: %s \n", pluginSpec.Yaml, pluginSpec.PluginName)
		err = kctl.Apply(config, pluginSpec, nil)
		if err != nil {
			return err, true
		}
	}
	return nil, false
}
