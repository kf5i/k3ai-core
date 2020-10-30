package cli

import (
	"fmt"

	"github.com/kf5i/k3ai-core/internal/plugins"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all plugins",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		config := newConfig()
		switch args[0] {
		case "plugins":
			{
				plugins, err := plugins.ContentList(pluginRepoURI)
				if err != nil {
					return err
				}
				for _, p := range plugins {
					fmt.Fprintln(config.Stdout(), p.Name)
				}

			}
		case "groups":
			{
				groups, err := plugins.ContentList(pluginsGroupRepoURI)
				if err != nil {
					return err
				}
				for _, g := range groups {
					fmt.Fprintln(config.Stdout(), g.Name)
				}
			}

		}

		return nil
	},
}
