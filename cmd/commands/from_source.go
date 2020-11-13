package commands

import (
	"fmt"
	"github.com/kf5i/k3ai-core/internal/k8s/kctl"
	"github.com/kf5i/k3ai-core/internal/plugins"
	"github.com/kf5i/k3ai-core/internal/shared"
)

const (
	ApplyOperation  = "apply"
	DeleteOperation = "delete"
)

func HandlePlugin(config kctl.Config, repo string, pluginName string, operation string) error {
	var plugin plugins.Plugin
	err := plugin.Encode(
		shared.NormalizeURL(repo, plugins.PluginDir, pluginName) + plugins.DefaultPluginFileName)
	if err != nil {
		return err
	}
	fmt.Printf("Plugin YAML content: %s, name: %s \n", plugin.Yaml, plugin.PluginName)
	switch operation {
	case ApplyOperation:
		{
			err = kctl.Apply(config, plugin, &kctl.CliWait{})
		}
	case DeleteOperation:
		err = kctl.Delete(config, plugin)
	}
	if err != nil {
		return err
	}
	return nil
}
