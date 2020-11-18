package infra

import (
	"fmt"
	"os"
	"os/exec"
)

func Kind(osFlavor string, infraSelection string) {
	// where are we? If windows we have to call wsl function if not proceed
	switch osFlavor {
	case "windows":
		infraKindWSL()
	case "linux":
		infraKindDefault()
	case "arm":
		infraARM()
	default:
		infraKindDefault()
	}
	// check if K3s is in the path
	// if k3s is not in the path download it
	// now we do install with the default flags

}

func infraKindWSL() {
	// we are in WSL so we cannot use the default installer
	cmd := exec.Command("wsl", "curl", "-Lo", "./kind https://kind.sigs.k8s.io/dl/v0.9.0/kind-linux-amd64", "^|", "bash", "-s", "--", "--wsl")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

}

func infraKindDefault() {
	// Let's download and install K3s the usual way
	cmd := exec.Command("/bin/sh", "-c", "curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.9.0/kind-linux-amd64; chmod +x ./kind; sudo mv ./kind /usr/local/bin; kind create cluster;")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func infraKindARM() {
	//ARM installation
}
