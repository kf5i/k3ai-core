package infra

import (
	"fmt"
	"github.com/enescakir/emoji"
	"github.com/kf5i/k3ai-core/internal/shared"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

// K0s check the OS flavor and provide an input to the subsequent functions
func K0s(osFlavor string, infraSelection string) {
	// where are we? If windows we have to call wsl function if not proceed
	switch osFlavor {
	case "windows":
		infraK0sWSL(osFlavor)
	case "linux":
		infraK0sDefault(osFlavor)
	case "arm":
		infraK0sARM()
	case "darwin":
		fmt.Println("Sorry K0s is not yet supported on your system")
	default:
		infraK0sDefault(osFlavor)
	}
	// check if K3s is in the path
	// if k3s is not in the path download it
	// now we do install with the default flags

}

func infraK0sWSL(osFlavor string) {
	// we are in WSL so we cannot use the default installer
	fmt.Printf("Hold on %v, we are going to install K0s %v\n", emoji.VulcanSalute, emoji.BuildingConstruction)
	time.Sleep(3 * time.Second)
	checkK3stest := true
	checkK3s := shared.CommandExists("k0s", osFlavor, checkK3stest)

	if checkK3s != true {
		cmd := exec.Command("bash", "-c", "curl -Lo ./k0s https://github.com/k0sproject/k0s/releases/download/v0.7.0/k0s-v0.7.0-amd64; chmod +x ./k0s; sudo mv ./k0s /usr/local/bin;mkdir -p ${HOME}/.k0s")
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout

		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		shared.CallClear()
		fmt.Printf("K0s %v! Now we are going to complete the setup %v, but first we need some help from you...\n", emoji.OkButton, emoji.Rocket)
		launchK0sFile(osFlavor)

	} else {
		fmt.Printf("K0s %v! Now we are going to complete the setup %v, but first we need some help from you...\n", emoji.OkButton, emoji.Rocket)
		launchK0sFile(osFlavor)
	}
}

func infraK0sDefault(osFlavor string) {
	// Let's download and install K3s the usual way
	fmt.Printf("Hold on %v, we are going to install K3s %v\n", emoji.VulcanSalute, emoji.BuildingConstruction)
	time.Sleep(3 * time.Second)
	checkK3stest := true
	checkK3s := shared.CommandExists("k0s", osFlavor, checkK3stest)

	if checkK3s != true {
		cmd := exec.Command("/bin/sh", "-c", "curl -Lo ./k0s https://github.com/k0sproject/k0s/releases/download/v0.7.0/k0s-v0.7.0-amd64; chmod +x ./k0s; sudo mv ./k0s /usr/local/bin;mkdir -p ${HOME}/.k0s")
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout

		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		shared.CallClear()
		fmt.Printf("K0s %v! Now we are going to complete the setup %v, but first we need some help from you...\n", emoji.OkButton, emoji.Rocket)
		launchK0sFile(osFlavor)

	} else {
		fmt.Printf("K0s %v! Now we are going to complete the setup %v, but first we need some help from you...\n", emoji.OkButton, emoji.Rocket)
		launchK0sFile(osFlavor)
	}
}

func infraK0sARM() {
	//ARM installation
}

func runK0sDefault(osFlavor string) {
	fmt.Println("file written successfully")
	if osFlavor == "windows" {
		cmd := exec.Command("bash", "-c", "sudo mv ./start.sh ${HOME}/.k0s/; chmod +x ${HOME}/.k0s/start.sh ;. ${HOME}/.k0s/start.sh")
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout

		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("K3ai installation complete %v%v%v!\n", emoji.PartyPopper, emoji.PartyPopper, emoji.PartyPopper)
		fmt.Printf("To use K3ai copy the following line. Once done type in your terminal: wsl press ctrl+c followed by Enter on your keyboard %v\n", emoji.MechanicalArm)
		fmt.Printf("%v  $HOME/.k0s/start.sh && export KUBECONFIG=/var/lib/k0s/pki/admin.conf\n", emoji.RightArrow)
		fmt.Printf("Thank you again for using K3ai, don't forget to check our docs at %v https://docs.k3ai.in\n", emoji.WorldMap)
	} else {
		cmd := exec.Command("/bin/bash", "-c", "sudo mv ./start.sh ${HOME}/.k0s/; chmod +x ${HOME}/.k0s/start.sh ;. /${HOME}/.k0s/start.sh")
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout

		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("K3ai installation complete %v%v%v!\n", emoji.PartyPopper, emoji.PartyPopper, emoji.PartyPopper)
		fmt.Printf("To use K3ai copy the following line: %v\n", emoji.MechanicalArm)
		fmt.Printf("%v  export KUBECONFIG=/var/lib/k0s/pki/admin.conf\n", emoji.RightArrow)
		fmt.Printf("Thank you again for using K3ai, don't forget to check our docs at %v https://docs.k3ai.in\n", emoji.WorldMap)

	}

}

func launchK0sFile(osFlavor string) {
	err := ioutil.WriteFile("start.sh", []byte("#!/bin/bash\n"), 0644)
	f, err := os.Create("start.sh")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	d := []string{"#!/bin/bash", "sudo whoami", "sudo k0s server -c ${HOME}/.k0s/k0s.yaml --enable-worker > /dev/null 2>&1 &"}

	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	runK0sDefault(osFlavor)
}
