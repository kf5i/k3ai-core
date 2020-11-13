package cli

import (
	"github.com/kf5i/k3ai-core/cmd/commands"
	"github.com/kf5i/k3ai-core/internal/k8s/kctl"
	"github.com/kf5i/k3ai-core/internal/plugins"
	"github.com/spf13/cobra"
)

func newDeleteCommand() *cobra.Command {
	var deleteCmd = &cobra.Command{
		Use:   "delete <plugin_name>",
		Short: "Delete a plugin or a plugin group",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			config := newConfig(cmd)

			group, _ := cmd.Flags().GetBool(plugins.GroupType)
			if group {
				return deleteGroup(config, args[0])

			}
			return deletePlugin(config, args[0])
		},
	}
	deleteCmd.Flags().BoolP(plugins.GroupType, "g", false, "Delete a plugin group")
	return deleteCmd
}
func deleteGroup(config kctl.Config, groupName string) error {
	var groups plugins.Groups
	pluginsGroupSpec, err := groups.Encode(repo+plugins.GroupsDir, groupName)
	if err != nil {
		return err
	}

	for _, group := range pluginsGroupSpec.Groups {
		for _, plugin := range group.Plugins {
			if plugin.Enabled {
				err := deletePlugin(config, plugin.Name)
				if err != nil {
					return err
				}
			}
		}

		for _, inlinePlugin := range group.InlinePlugins {
			err = kctl.Delete(config, inlinePlugin)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func deletePlugin(config kctl.Config, pluginName string) error {
	return commands.HandlePlugin(config, repo, pluginName, commands.DeleteOperation)
}
