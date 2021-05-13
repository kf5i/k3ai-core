package kctl

import (
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/kf5i/k3ai-core/internal/plugins"
)

const (
	kubectl    = "kubectl"
	helm       = "helm"
	apply      = "apply"
	delete     = "delete"
	create     = "create"
	helmApply  = "install"
	helmDelete = "delete"
)

func pause() {
	time.Sleep(2 * time.Second)
}

// Apply adds/updates the plugin in a k3s/k8s cluster
func Apply(config Config, plugin plugins.Plugin, evt Wait) error {
	_ = createNameSpace(config, plugin.Namespace)
	pause()
	for _, yamlSpec := range plugin.Yaml {
		var err error
		if decodeType(yamlSpec.Type) == "helm" {
			command, args := prepareCommand(config, helm, helmApply,
				yamlSpec.URL)
			err = execute(config, command, args...)
		} else if decodeType(yamlSpec.Type) == "container" {
			command, args := prepareCommand(config, apply,
				decodeType(yamlSpec.Type), yamlSpec.URL, "-n", plugin.Namespace)
			err = execute(config, command, args...)
		} else {
			command, args := prepareCommand(config, apply,
				decodeType(yamlSpec.Type), yamlSpec.URL, "-n", plugin.Namespace)
			err = execute(config, command, args...)
		}

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
	var err error
	for i := len(plugin.Yaml) - 1; i >= 0; i-- {
		yamlSpec := plugin.Yaml[i]
		if decodeType(yamlSpec.Type) == "helm" {
			command, args := prepareCommand(config, helm, helmDelete,
				yamlSpec.URL, "--namespace", plugin.Namespace)
			err = execute(config, command, args...)
		} else if decodeType(yamlSpec.Type) == "container" {

		} else {
			command, args := prepareCommand(config, delete,
				decodeType(yamlSpec.Type), yamlSpec.URL, "-n", plugin.Namespace)
			err = execute(config, command, args...)
		}
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
	if commandType == plugins.CommandHelm {
		return "helm"
	}
	return "-f"
}

func execute(config Config, command string, args ...string) error {
	if command == helm {
		var installCmd []string
		copy(args[0:], args[0+1:]) // Shift a[i+1:] left one index.
		args[len(args)-1] = ""
		args = args[:len(args)-1]
		installCmd = append(installCmd, args[0])
		if args[0] == "delete" {
			installCmd = append(installCmd, strings.Split(args[1], " ")...)
			cmd := exec.Command(command, installCmd[0], installCmd[1], args[2], args[3])
			log.Print(cmd)
			cmd.Stdout = config.Stdout()
			cmd.Stderr = config.Stderr()
			return cmd.Run()
		}
		installCmd = append(installCmd, strings.Split(args[1], " ")...)
		cmd := exec.Command(command, installCmd...)
		log.Print(cmd)
		cmd.Stdout = config.Stdout()
		cmd.Stderr = config.Stderr()
		return cmd.Run()
	}
	cmd := exec.Command(command, args...)
	cmd.Stdout = config.Stdout()
	cmd.Stderr = config.Stderr()
	return cmd.Run()
}

func prepareCommand(config Config, args ...string) (string, []string) {
	if args[0] == helm {
		command := helm
		return command, args
	}
	command := kubectl
	return command, args
}
