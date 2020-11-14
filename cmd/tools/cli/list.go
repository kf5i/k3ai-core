package cli

import (
	"fmt"

	"github.com/kf5i/k3ai-core/internal/plugins"
	"github.com/spf13/cobra"
)

func newListCommand() *cobra.Command {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all plugins or plugin groups",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			config := newConfig(cmd)
			group, _ := cmd.Flags().GetBool(plugins.GroupType)
			if group {
				groups, err := plugins.ContentList(repo + plugins.GroupsDir)
				if err != nil {
					return err
				}
				for _, g := range groups {
					fmt.Fprintln(config.Stdout(), g.Name)
				}
				return nil
			}
			var pls plugins.Plugins
			err := pls.List(repo + plugins.PluginDir)
			if err != nil {
				return err
			}

			PrintFormat("Name", "Description")
			for _, p := range pls.Items {
				PrintFormat(p.PluginName, p.PluginDescription)

			}

			return nil
		},
	}
	listCmd.Flags().BoolP(plugins.GroupType, "g", false, "List the plugin groups")
	return listCmd
}
