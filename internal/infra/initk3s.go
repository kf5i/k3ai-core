package infra

/*Author: Alessandro Festa
Infra package allow a user to install a local cluster based on:
K3s
K0s
Kind
*/
import (
	"fmt"
	"os"
	"os/exec"
)

// K3s check the OS flavor and provide an input to the subsequent functions
func K3s(osFlavor string, infraSelection string) {
	// where are we? If windows we have to call wsl function if not proceed
	switch osFlavor {
	case "windows":
		infraWSL()
	case "linux":
		infraDefault()
	case "arm":
		infraARM()
	default:
		infraDefault()
	}
	// check if K3s is in the path
	// if k3s is not in the path download it
	// now we do install with the default flags

}

func infraWSL() {
	// we are in WSL so we cannot use the default installer
	cmd := exec.Command("wsl", "curl", "-sfL", "https://get.k3ai.in", "^|", "bash", "-s", "--", "--wsl")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

}

func infraDefault() {
	// Let's download and install K3s the usual way
	cmd := exec.Command("/bin/sh", "-c", "curl -sfL https://get.k3s.io | K3S_KUBECONFIG_MODE=644 sh -s -; export KUBECONFIG=/etc/rancher/k3s/k3s.yaml")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func infraARM() {
	//ARM installation
}
