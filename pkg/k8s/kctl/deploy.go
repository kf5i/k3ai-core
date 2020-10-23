package kctl

import (
	"bytes"
	"log"
	"os/exec"

	"github.com/kf5i/k3ai-core/pkg/plugins"
)

const (
	k3sExec = "k3s"
	kubectl = "kubectl"
	apply   = "apply"
	delete  = "delete"
)

type (
	Wait interface {
		Process(labels []string)
	}
)

func Apply(plugin plugins.PluginSpec, evt Wait) error {
	err := handleYaml(apply, plugin)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if evt != nil {
		evt.Process(plugin.Yaml)
	}
	return nil
}

func Delete(plugin plugins.PluginSpec) error {
	return handleYaml(delete, plugin)
}

func handleYaml(command string, plugin plugins.PluginSpec) error {
	for _, fileYaml := range plugin.Yaml {
		cmd := exec.Command(k3sExec, kubectl, command, "-f", fileYaml)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Print(err)
			return err
		}
		log.Printf(" %q\n", out.String())
	}
	return nil
}
