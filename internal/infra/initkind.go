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
	"github.com/enescakir/emoji"
	"github.com/kf5i/k3ai-core/internal/shared"
	"os"
	"os/exec"
	"time"
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
	fmt.Printf("Hold on %v, we are going to install Kind %v\n", emoji.VulcanSalute, emoji.BuildingConstruction)
	time.Sleep(3 * time.Second)
	checkK3stest := true
	checkK3s := shared.CommandExists("kind", osFlavor, checkK3stest)

	if checkK3s != true {
		cmd := exec.Command("powershell", "curl.exe -Lo kind.exe https://kind.sigs.k8s.io/dl/v0.9.0/kind-windows-amd64; Move-Item ./kind.exe -Destination C:/Windows/System32/ -force ; kind create cluster")
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout

		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Kind installation complete %v%v%v!\n", emoji.PartyPopper, emoji.PartyPopper, emoji.PartyPopper)
		fmt.Printf("To use K3ai follow these steps:\n")
		fmt.Printf("If you are on Windows/WSL type in your terminal: wsl and get start use K3ai %v\n", emoji.MechanicalArm)
		fmt.Printf("If you are on macOS copy and paste the line below %v\n", emoji.MechanicalArm)
		fmt.Printf("%v  export KUBECONFIG=/var/lib/k0s/pki/admin.conf\n", emoji.RightArrow)
		fmt.Printf("Thank you again for using K3ai, don't forget to check our docs at %v https://docs.k3ai.in\n", emoji.WorldMap)

	} else {
		fmt.Printf("Kind installation complete %v%v%v!\n", emoji.PartyPopper, emoji.PartyPopper, emoji.PartyPopper)
		fmt.Printf("To use K3ai follow these steps:\n")
		fmt.Printf("If you are on Windows/WSL type in your terminal: wsl and get start use K3ai %v\n", emoji.MechanicalArm)
		fmt.Printf("If you are on macOS copy and paste the line below %v\n", emoji.MechanicalArm)
		fmt.Printf("%v  export KUBECONFIG=/var/lib/k0s/pki/admin.conf\n", emoji.RightArrow)
		fmt.Printf("Thank you again for using K3ai, don't forget to check our docs at %v https://docs.k3ai.in\n", emoji.WorldMap)
	}
}

func infraKindDefault(osFlavor string) {
	// Let's check if we are in MacOs or in Linux and for Linux if we are within WSL)
	fmt.Printf("Hold on %v, we are going to install Kind %v\n", emoji.VulcanSalute, emoji.BuildingConstruction)
	time.Sleep(3 * time.Second)
	checkK3stest := true
	checkK3s := shared.CommandExists("kind", osFlavor, checkK3stest)

	if checkK3s != true {
		var cmd *exec.Cmd
		if osFlavor == "darwin" {
			cmd = exec.Command("/bin/sh", "-c", "curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.9.0/kind-darwin-amd64; chmod +x ./kind; sudo mv ./kind /usr/local/bin")

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
		fmt.Printf("Kind installation complete %v%v%v!\n", emoji.PartyPopper, emoji.PartyPopper, emoji.PartyPopper)
		fmt.Printf("To use K3ai follow these steps:\n")
		fmt.Printf("If you are on Windows/WSL type in your terminal: wsl and get start use K3ai %v\n", emoji.MechanicalArm)
		fmt.Printf("If you are on macOS copy and paste the line below %v\n", emoji.MechanicalArm)
		fmt.Printf("%v  export KUBECONFIG=/var/lib/k0s/pki/admin.conf\n", emoji.RightArrow)
		fmt.Printf("Thank you again for using K3ai, don't forget to check our docs at %v https://docs.k3ai.in\n", emoji.WorldMap)
	}
}

func infraKindARM() {
	//ARM installation
}
