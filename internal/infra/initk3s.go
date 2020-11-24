package infra

/*Author: Alessandro Festa
Infra package allow a user to install a local cluster based on:
K3s
K3s
Kind
*/
import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
	"github.com/enescakir/emoji"
	"github.com/kf5i/k3ai-core/internal/shared"
)

// K3s check the OS flavor and provide an input to the subsequent functions
func K3s(osFlavor string, infraSelection string) {
	// where are we? If windows we have to call wsl function if not proceed
	switch osFlavor {
	case "windows":
		infraK3sWSL(osFlavor)
	case "linux":
		infraK3sDefault(osFlavor)
	case "arm":
		infraK3sARM()
	case "darwin":
		fmt.Println("Sorry K3s is not yet supported on your system")
	default:
		infraK3sDefault(osFlavor)
	}
	// check if K3s is in the path
	// if k3s is not in the path download it
	// now we do install with the default flags

}

func infraK3sWSL(osFlavor string) {
	// we are in WSL so we cannot use the default installer
	fmt.Printf("Hold on %v, we are going to install K3s %v\n",emoji.VulcanSalute,emoji.BuildingConstruction)
	time.Sleep(3 * time.Second)
	checkK3stest := true
	checkK3s := shared.CommandExists("k3s", osFlavor, checkK3stest)

	if checkK3s != true {
		cmd := exec.Command("bash", "-c", "curl -Lo ./k3s https://github.com/rancher/k3s/releases/download/v1.19.4%2Bk3s1/k3s; chmod +x ./k3s; sudo mv ./k3s /usr/local/bin;mkdir -p ${HOME}/.k3s")
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
	
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		shared.CallClear()
		fmt.Printf("K3s %v! Now we are going to complete the setup %v, but first we need some help from you...\n", emoji.OkButton, emoji.Rocket)
		launchK3sFile(osFlavor)

	} else {
		fmt.Printf("K3s %v! Now we are going to complete the setup %v, but first we need some help from you...\n", emoji.OkButton, emoji.Rocket)
		launchK3sFile(osFlavor)
	}
}

func infraK3sDefault(osFlavor string) {
	// Let's download and install K3s  but first check if we are inside WSL session
	var cmd *exec.Cmd
	fmt.Printf("Hold on %v, we are going to install K3s %v\n",emoji.VulcanSalute,emoji.BuildingConstruction)
	time.Sleep(3 * time.Second)
	checkK3stest := true
	checkK3s := shared.CommandExists("k3s", osFlavor, checkK3stest)
	if checkK3s != true {
		if os.Getenv("WSL_DISTRO_NAME") != "" {
			cmd = exec.Command("/bin/sh", "-c", "curl -Lo ./k3s https://github.com/rancher/k3s/releases/download/v1.19.4%2Bk3s1/k3s ; chmod +x ./k3s; sudo mv ./k3s /usr/local/bin;mkdir -p ${HOME}/.k3s")
		} else {
			cmd = exec.Command("/bin/sh", "-c", "curl -sfL https://get.k3s.io | K3S_KUBECONFIG_MODE=644 sh -s -")
		}
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
	
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		shared.CallClear()
		fmt.Printf("K3s %v! Now we are going to complete the setup %v, but first we need some help from you...\n", emoji.OkButton, emoji.Rocket)
		launchK3sFile(osFlavor)
	} else {
		fmt.Printf("K3s %v! Now we are going to complete the setup %v, but first we need some help from you...\n", emoji.OkButton, emoji.Rocket)
		launchK3sFile(osFlavor)
	}

}

func infraK3sARM() {
	//ARM installation
}

func runK3sDefault(osFlavor string) {

	if osFlavor == "windows" {
		cmd := exec.Command("bash", "-c", "sudo mv ./start.sh ${HOME}/.k3s/; chmod +x ${HOME}/.k3s/start.sh ;. /${HOME}/.k3s/start.sh")
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout

		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("K3ai installation complete %v%v%v!\n", emoji.PartyPopper,emoji.PartyPopper,emoji.PartyPopper)
		fmt.Printf("To use K3ai copy the following line. Once done type in your terminal: wsl press ctrl+c followed by Enter on your keyboard %v\n", emoji.MechanicalArm)
		fmt.Printf("%v  export KUBECONFIG=/etc/rancher/k3s/k3s.yaml\n",emoji.RightArrow)
		fmt.Printf("Thank you again for using K3ai, don't forget to check our docs at %v https://docs.k3ai.in\n", emoji.WorldMap)
	} else {
		cmd := exec.Command("/bin/bash", "-c", "sudo mv ./start.sh ${HOME}/.k3s/; chmod +x ${HOME}/.k3s/start.sh ;. /${HOME}/.k3s/start.sh")
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout

		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("K3ai installation complete %v%v%v!\n", emoji.PartyPopper,emoji.PartyPopper,emoji.PartyPopper)
		fmt.Printf("To use K3ai copy the following line: %v\n", emoji.MechanicalArm)
		fmt.Printf("%v  export KUBECONFIG=/etc/rancher/k3s/k3s.yaml\n",emoji.RightArrow)
		fmt.Printf("Thank you again for using K3ai, don't forget to check our docs at %v https://docs.k3ai.in\n", emoji.WorldMap)

	}

}

func launchK3sFile(osFlavor string) {
	err := ioutil.WriteFile("start.sh", []byte("#!/bin/bash\n"), 0644)
	f, err := os.Create("start.sh")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	d := []string{"#!/bin/bash", "sudo nohup k3s server --write-kubeconfig-mode 644 > /dev/null 2>&1 &"}

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
	runK3sDefault(osFlavor)
}

