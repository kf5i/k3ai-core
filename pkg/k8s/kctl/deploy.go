package kctl

import (
	"bytes"
	"log"
	"os/exec"

	"github.com/kf5i/k3ai-core/pkg/plugins"
)

const K3sExec = "k3s"
const Kubectl = "kubectl"
const Apply = "apply"
const Delete = "delete"

type (
	Wait interface {
		Process(labels []string)
	}
)

func ApplyFiles(plugin plugins.PluginSpec, evt Wait) error {
	err := handleFiles(Apply, plugin)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if evt != nil {
		evt.Process(plugin.Files)
	}
	return nil
}

func DeleteFiles(plugin plugins.PluginSpec) error {
	return handleFiles(Delete, plugin)
}

func handleFiles(command string, plugin plugins.PluginSpec) error {
	for _, fileYaml := range plugin.Files {
		cmd := exec.Command(K3sExec, Kubectl, command, "-f", fileYaml)
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
