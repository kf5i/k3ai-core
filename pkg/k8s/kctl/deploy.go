package kctl

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

const K3sExec = "k3s"
const Kublectl = "kubectl"
const Apply = "apply"
const Delete = "delete"

// const K3S_APPLY_KUSTOMIZE = "k3 kubectl apply -k "

func ApplyFiles(filenames []string) error {
	return handleFiles(Apply, filenames)
}

func DeleteFiles(filenames []string) error {
	return handleFiles(Delete, filenames)
}

func handleFiles(command string, filenames []string) error {
	for _, fileYaml := range filenames {
		fmt.Println(fileYaml)
		cmd := exec.Command(K3sExec, Kublectl, command, "-f", fileYaml)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(" %q\n", out.String())
	}
	return nil
}
