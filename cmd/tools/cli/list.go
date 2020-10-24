package cli

import (
	"fmt"

	"github.com/kf5i/k3ai-core/internal/plugins"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all plugins",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		config := newConfig()
		pluginList, err := plugins.GetPluginList()
		if err != nil {
			return err
		}
		for _, p := range *pluginList {
			fmt.Fprintln(config.Stdout(), p.Name)
		}
		return nil
	},
}
