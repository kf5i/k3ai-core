package kctl

import (
	"log"
	"os/exec"
	"time"

	"github.com/kf5i/k3ai-core/internal/plugins"
)

const (
	kubectl = "kubectl"
	apply   = "apply"
	delete  = "delete"
	create  = "create"
)

func pause() {
	time.Sleep(2 * time.Second)
}

// Apply adds/updates the plugin in a k3s/k8s cluster
func Apply(config Config, plugin plugins.Plugin, evt Wait) error {
	_ = createNameSpace(config, plugin.Namespace)
	pause()
	for _, yamlSpec := range plugin.Yaml {
		command, args := prepareCommand(config, apply,
			decodeType(yamlSpec.Type), yamlSpec.URL, "-n", plugin.Namespace)
		err := execute(config, command, args...)
		if err != nil {
			log.Printf("Error during create: %s\n", err.Error())
		}
		pause()
	}

	if evt != nil {
		err := evt.Process(config, plugin.Namespace, plugin.Labels)
		if err != nil {
			log.Printf("Error during wait: %s\n", err.Error())
		}
	}

	if plugin.PostInstall.Command != "" {

		err := execute(config, "sh", "-c", plugin.PostInstall.Command)
		if err != nil {
			log.Printf("Error during post installation: %s\n", err.Error())
		}
	}
	return nil
}

// Delete removes the plugin from the cluster
func Delete(config Config, plugin plugins.Plugin) error {
	for i := len(plugin.Yaml) - 1; i >= 0; i-- {
		yamlSpec := plugin.Yaml[i]
		command, args := prepareCommand(config, delete,
			decodeType(yamlSpec.Type), yamlSpec.URL, "-n", plugin.Namespace)
		err := execute(config, command, args...)
		if err != nil {
			log.Printf("Error during delete: %s\n", err.Error())

		}
		pause()
	}

	return nil
}

func decodeType(commandType string) string {
	if commandType == plugins.CommandKustomize {
		return "-k"
	}
	return "-f"
}

func execute(config Config, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = config.Stdout()
	cmd.Stderr = config.Stderr()
	return cmd.Run()
}

func prepareCommand(config Config, args ...string) (string, []string) {
	command := kubectl
	return command, args
}
