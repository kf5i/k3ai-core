package cli

import (
	"fmt"

	"github.com/kf5i/k3ai-core/internal/plugins"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all plugins or plugin groups",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		config := newConfig()
		group, _ := cmd.Flags().GetBool(plugins.GroupType)
		if group {
			groups, err := plugins.ContentList(pluginsGroupRepoURI)
			if err != nil {
				return err
			}
			for _, g := range groups {
				fmt.Fprintln(config.Stdout(), g.Name)
			}
			return nil
		}
		plugins, err := plugins.ContentList(pluginRepoURI)
		if err != nil {
			return err
		}
		for _, p := range plugins {
			fmt.Fprintln(config.Stdout(), p.Name)
		}

		return nil
	},
}
