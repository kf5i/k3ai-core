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
			group, _ := cmd.Flags().GetBool(plugins.GroupType)
			if group {

				var cache plugins.Cache
				err := cache.Encode(repo, "cache_groups.yaml")
				if err != nil {
					fmt.Printf("Can't load cache:%s will use remote\n", err)
					var grs plugins.Groups
					err := grs.List(repo + plugins.GroupsDir)
					if err != nil {
						return err
					}
					PrintFormat("Name", "Description")
					for _, p := range grs.Items {
						PrintFormat(p.GroupName, p.GroupDescription)
					}
					return nil
				}
				PrintFormat("Name", "Description")
				for _, p := range cache.Items {
					PrintFormat(p.Name, p.Description)
				}

				return nil
			}

			var cache plugins.Cache
			err := cache.Encode(repo, "cache_plugins.yaml")
			if err != nil {
				fmt.Printf("Can't load cache:%s will use remote\n", err)

				var pls plugins.Plugins
				err = pls.List(repo + plugins.PluginDir)
				if err != nil {
					return err
				}

				PrintFormat("Name", "Description")
				for _, p := range pls.Items {
					PrintFormat(p.PluginName, p.PluginDescription)
				}
				return nil
			}
			PrintFormat("Name", "Description")
			for _, p := range cache.Items {
				PrintFormat(p.Name, p.Description)
			}

			return nil
		},
	}
	listCmd.Flags().BoolP(plugins.GroupType, "g", false, "List the plugin groups")
	return listCmd
}
