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
	create  = "create"
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
		evt.Process(plugin.Labels)
	}
	return nil
}

func Delete(config Config, plugin plugins.PluginSpec) error {
	return handleYaml(config, delete, plugin)
}

func decodeType(commandType string) string {
	if commandType == plugins.CommandKustomize {
		return "-k"
	}
	return "-f"
}

func handleYaml(config Config, command string, plugin plugins.PluginSpec) error {
	for _, yamlSpec := range plugin.Yaml {
		if command == apply {
			_ = createNameSpace(config, yamlSpec.NameSpace)
		}
		err := execute(config, k3sExec, kubectl, command,
			decodeType(yamlSpec.Type), yamlSpec.Url, "-n", yamlSpec.NameSpace)
		if err != nil {
			return err
		}
		if command == delete {
			_ = deleteNameSpace(config, yamlSpec.NameSpace)
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
