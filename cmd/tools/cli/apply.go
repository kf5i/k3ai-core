package cli

import (
	"github.com/kf5i/k3ai-core/cmd/commands"
	"github.com/kf5i/k3ai-core/internal/k8s/kctl"
	"github.com/kf5i/k3ai-core/internal/plugins"
	"github.com/spf13/cobra"
)

func newApplyCommand() *cobra.Command {
	var applyCmd = &cobra.Command{
		Use:   "apply <plugin_name>",
		Short: "Apply a plugin or a plugin group",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			config := newConfig(cmd)
			group, _ := cmd.Flags().GetBool(plugins.GroupType)
			if group {
				return applyGroup(config, args[0])
			}
			return applyPlugin(config, args[0])
		},
	}
	applyCmd.Flags().BoolP(plugins.GroupType, "g", false, "Apply a plugin group")
	return applyCmd
}

func applyGroup(config kctl.Config, groupName string) error {
	var groups plugins.Groups
	pluginsGroupSpec, err := groups.Encode(repo+plugins.GroupsDir, groupName)
	if err != nil {
		return err
	}

	for _, group := range pluginsGroupSpec.Groups {
		for _, plugin := range group.Plugins {
			if plugin.Enabled {
				err := applyPlugin(config, plugin.Name)
				if err != nil {
					return err
				}
			}
		}

		for _, inlinePlugin := range group.InlinePlugins {
			err = kctl.Apply(config, inlinePlugin, &kctl.CliWait{})
			if err != nil {
				return err
			}
		}
	}

	return nil

}

func applyPlugin(config kctl.Config, pluginName string) error {
	return commands.HandlePlugin(config, repo, pluginName, commands.ApplyOperation)
}
