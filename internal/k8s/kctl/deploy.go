package kctl

import (
	"os/exec"

	"github.com/kf5i/k3ai-core/internal/plugins"
)

const (
	k3sExec = "k3s"
	kubectl = "kubectl"
	apply   = "apply"
	delete  = "delete"
)

type Wait interface {
	Process(labels []string)
}

func Apply(config Config, plugin plugins.PluginSpec, evt Wait) error {
	err := handleYaml(config, apply, plugin)
	if err != nil {
		return err
	}
	if evt != nil {
		evt.Process(plugin.Yaml)
	}
	return nil
}

func Delete(config Config, plugin plugins.PluginSpec) error {
	return handleYaml(config, delete, plugin)
}

func handleYaml(config Config, command string, plugin plugins.PluginSpec) error {
	for _, fileYaml := range plugin.Yaml {
		err := execute(config, k3sExec, kubectl, command, "-f", fileYaml)
		if err != nil {
			return err
		}
	}
	return nil
}

func execute(config Config, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = config.Stdout()
	cmd.Stderr = config.Stderr()
	return cmd.Run()
}
