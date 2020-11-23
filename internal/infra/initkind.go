package infra

/*Author: Alessandro Festa
Infra package allow a user to install a local cluster based on:
K3s
Kind
Kind
This is the Kind installation
*/

//cmd := exec.Command("wsl", "curl", "-Lo", "./kind https://kind.sigs.k8s.io/dl/v0.9.0/kind-linux-amd64", "^|", "bash", "-s", "--", "--wsl")

import (
	"fmt"
	"os"
	"os/exec"
)

// Kind check the OS flavor and provide an input to the subsequent functions
func Kind(osFlavor string, infraSelection string) {
	// where are we? If windows we have to call wsl function if not proceed
	switch osFlavor {
	case "windows":
		infraKindWSL(osFlavor)
	case "linux":
		infraKindDefault(osFlavor)
	case "arm":
		infraKindARM()
	case "darwin":
		infraKindDefault(osFlavor)
	default:
		infraKindDefault(osFlavor)
	}
	// check if K3s is in the path
	// if k3s is not in the path download it
	// now we do install with the default flags

}

func infraKindWSL(osFlavor string) {
	// we are in WSL so we cannot use the default installer
	cmd := exec.Command("powershell", "curl.exe -Lo kind-windows-amd64.exe https://kind.sigs.k8s.io/dl/v0.9.0/kind-windows-amd64; New-Item -ItemType directory -Path ${HOME}/.kind ; Move-Item ./kind-windows-amd64.exe ${HOME}/.kind/ -force ; ./${HOME}/.kind/kind-windows-amd64.exe create cluster")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func infraKindDefault(osFlavor string) {
	// Let's check if we are in MacOs or in Linux and for Linux if we are within WSL)
	var cmd *exec.Cmd
	if osFlavor == "darwin" {
		cmd = exec.Command("/bin/sh", "-c", "curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.9.0/kind-darwin-amd64; chmod +x ./kind; sudo mv ./kind /usr/local/bin; echo 'Copy and Paste the following lines to run kind'; echo 'export PATH=/usr/local/bin:$PATH'; echo 'kind create cluster'")
	} else if os.Getenv("WSL_DISTRO_NAME") != "" {
		cmd = exec.Command("/bin/sh", "-c", "curl -Lo kind  https://kind.sigs.k8s.io/dl/v0.9.0/kind-linux-amd64; chmod +x ./kind; sudo mv ./kind /usr/local/bin; kind create cluster")
	} else {
		cmd = exec.Command("/bin/sh", "-c", "curl -Lo kind  https://kind.sigs.k8s.io/dl/v0.9.0/kind-linux-amd64; chmod +x ./kind; sudo mv ./kind /usr/local/bin; kind create cluster")
	}

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
