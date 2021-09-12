package local

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/enescakir/emoji"
	"github.com/kf5i/k3ai-core/cmd/commands"
	"github.com/kf5i/k3ai-core/internal/k8s/kctl"
	"github.com/kf5i/k3ai-core/internal/shared"
)

var cfg shared.Config

// Init start the Deployment for local configurations
func Init(kctlConfig kctl.Config, defaultRepo string, data shared.TargetCustoms) {
	finished := make(chan bool)
	fmt.Printf("%v	Checking requirements for local deployment...\n", emoji.CheckBoxWithCheck)
	go localDeployment(kctlConfig, defaultRepo, data, finished)
	<-finished
	fmt.Printf("%v	Local deployment completed, have fun with k3ai!\n", emoji.PartyPopper)
	if data.Type == "k3s" && data.ClusterDeployment == "local" {
		fmt.Printf("\n")
		time.Sleep(time.Second)
		fmt.Printf("To use K3ai copy the following line: %v\n", emoji.MechanicalArm)
		fmt.Printf("%v  export KUBECONFIG=%s\n", emoji.RightArrow, data.Config)
		fmt.Printf("Thank you again for using K3ai, don't forget to check our docs at %v https://docs.k3ai.in\n", emoji.WorldMap)
	} else if data.Type == "k0s" && data.ClusterDeployment == "local" {
		fmt.Printf("\n")
		time.Sleep(time.Second)
		fmt.Printf("To use K3ai copy the following line: %v\n", emoji.MechanicalArm)
		fmt.Printf("%v  export KUBECONFIG=%s\n", emoji.RightArrow, data.Config)
		fmt.Printf("Thank you again for using K3ai, don't forget to check our docs at %v https://docs.k3ai.in\n", emoji.WorldMap)
	}
}

func localDeployment(kctlConfig kctl.Config, defaultRepo string, data shared.TargetCustoms, finished chan bool) {
	fmt.Printf("%v	Installing infrastructure for local deployment...\n", emoji.CheckBoxWithCheck)
	fmt.Printf("	Selected deployment type: %s\n", data.Type)
	var cmd *exec.Cmd
	var osFlavor string
	osFlavor = strings.ToLower(runtime.GOOS)
	//kind deployment
	if data.Type == "kind" {
		if osFlavor == "darwin" {
			cmd = exec.Command("/bin/sh", "-c", "curl -Lo $HOME/kind "+data.Spec.Mac+"; chmod +x $HOME/kind; sudo mv $HOME/kind /usr/local/bin; ./usr/local/bin/"+data.ClusterStart+" --name "+data.ClusterName)
		} else if (osFlavor == "linux") && (os.Getenv("WSL_DISTRO_NAME") != "") {
			cmd = exec.Command("/bin/sh", "-c", "curl -Lo $HOME/kind "+data.Spec.Wsl+"; chmod +x $HOME/kind; sudo mv $HOME/kind /usr/local/bin;"+data.ClusterStart+" --name "+data.ClusterName)
		} else if (osFlavor == "linux") && (os.Getenv("WSL_DISTRO_NAME") == "") {
			cmd = exec.Command("/bin/sh", "-c", "curl -Lo $HOME/kind  "+data.Spec.Linux+"; chmod +x $HOME/kind; sudo mv $HOME/kind /usr/local/bin;"+data.ClusterStart+" --name "+data.ClusterName)
		} else {
			elevateCmd := "Start-Process powershell -ArgumentList 'Move-Item -Path ${HOME}/kind.exe -Destination C:/Windows/System32/ -force' -verb runAs"
			cmd = exec.Command("powershell", "curl.exe -Lo $HOME/kind.exe "+data.Spec.Windows+";"+elevateCmd+";powershell "+data.ClusterStart+" --name "+data.ClusterName)
		}
	}
	// k3s deployment, since in WSL and macOS is not currently possible to use the standard installer
	// we will inform the user at the end that k3d is a better option
	if data.Type == "k3s" {
		if osFlavor == "darwin" {
			cmd = exec.Command("/bin/sh", "-c", "curl -Lo $HOME/k3s "+data.Spec.Mac+"; chmod +x $HOME/k3s; sudo mv $HOME/k3s /usr/local/bin; ./usr/local/bin/"+data.ClusterStart)
		} else if (osFlavor == "linux") && (os.Getenv("WSL_DISTRO_NAME") != "") {
			cmd = exec.Command("/bin/sh", "-c", "curl -Lo $HOME/k3s "+data.Spec.Wsl+"; chmod +x $HOME/k3s; sudo mv $HOME/k3s /usr/local/bin;"+data.ClusterStart)
			shared.LaunchWSLFile(data.ClusterStart, data.Type)

		} else if (osFlavor == "linux") && (os.Getenv("WSL_DISTRO_NAME") == "") {
			cmd = exec.Command("/bin/sh", "-c", "curl -Lo $HOME/k3s  "+data.Spec.Linux+"; chmod +x $HOME/k3s; sudo mv $HOME/k3s /usr/local/bin;"+data.ClusterStart)
		} else {
			//since k3s cannot run natively on Windows we stop here and inform the user
			fmt.Printf("%v	Ops sorry %s is not yet supported on this OS...\n", emoji.StopSign, data.Type)
			time.Sleep(time.Second)
			os.Exit(9)

		}

	}
	// k0s deployment, since in WSL and macOS is not currently possible to use the standard installer
	// we will inform the user that some manual steps are required
	if data.Type == "k0s" {
		if osFlavor == "darwin" {
			cmd = exec.Command("/bin/sh", "-c", "curl -Lo $HOME/k0s "+data.Spec.Mac+"; chmod +x $HOME/k0s; sudo mv $HOME/k0s /usr/local/bin; ./usr/local/bin/"+data.ClusterStart)
		} else if (osFlavor == "linux") && (os.Getenv("WSL_DISTRO_NAME") != "") {
			cmd = exec.Command("/bin/sh", "-c", "curl -Lo $HOME/k0s "+data.Spec.Wsl+"; chmod +x $HOME/k0s; sudo mv $HOME/k0s /usr/local/bin;"+data.ClusterStart)
			shared.LaunchWSLFile(data.ClusterStart, data.Type)

		} else if (osFlavor == "linux") && (os.Getenv("WSL_DISTRO_NAME") == "") {
			cmd = exec.Command("/bin/sh", "-c", "curl -Lo $HOME/k0s  "+data.Spec.Linux+"; chmod +x $HOME/k0s; sudo mv $HOME/k0s /usr/local/bin;"+data.ClusterStart)
		} else {
			//since k3s cannot run natively on Windows we stop here and inform the user
			fmt.Printf("%v	Ops sorry %s is not yet supported on this OS...\n", emoji.StopSign, data.Type)
			time.Sleep(time.Second)
			os.Exit(9)

		}

	}
	//k3d deployment
	if data.Type == "k3d" {
		if osFlavor == "darwin" {
			cmd = exec.Command("/bin/sh", "-c", "curl -Lo $HOME/k3d "+data.Spec.Mac+"; chmod +x $HOME/k3d; sudo mv $HOME/k3d /usr/local/bin; ./usr/local/bin/"+data.ClusterStart+" "+data.ClusterName+" --update-default-kubeconfig --switch-context")
		} else if (osFlavor == "linux") && (os.Getenv("WSL_DISTRO_NAME") != "") {
			cmd = exec.Command("/bin/sh", "-c", "curl -Lo $HOME/k3d "+data.Spec.Wsl+"; chmod +x $HOME/k3d; sudo mv $HOME/k3d /usr/local/bin;"+data.ClusterStart+" "+data.ClusterName+" --update-default-kubeconfig --switch-context")
		} else if (osFlavor == "linux") && (os.Getenv("WSL_DISTRO_NAME") == "") {
			cmd = exec.Command("/bin/sh", "-c", "curl -Lo $HOME/k3d  "+data.Spec.Linux+"; chmod +x $HOME/k3d; sudo mv $HOME/k3d /usr/local/bin;"+data.ClusterStart+" "+data.ClusterName+" --update-default-kubeconfig --switch-context")
		} else {
			elevateCmd := "Start-Process powershell -ArgumentList 'Move-Item -Path ${HOME}/k3d.exe -Destination C:/Windows/System32/ -force' -verb runAs"
			cmd = exec.Command("powershell", "curl.exe -Lo $HOME/k3d.exe "+data.Spec.Windows+";"+elevateCmd+";powershell "+data.ClusterStart+" "+data.ClusterName+" --update-default-kubeconfig --switch-context")
		}
	}

	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = nil

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second)
	fmt.Printf("%v	Infrastructure ready, proceeding to plugins installation (if any)...\n", emoji.CheckBoxWithCheck)
	time.Sleep(time.Second)
	localAppDeployment(kctlConfig, defaultRepo, data)
	finished <- true
}

func localAppDeployment(kctlConfig kctl.Config, defaultRepo string, data shared.TargetCustoms) {
	fmt.Printf("%v	Add Plugins to local deployment...\n", emoji.CheckBoxWithCheck)

	errors := false
	for _, item := range data.Plugins {
		pluginName := item.Name
		if pluginName == "" {
			continue
		}
		repo := item.Repo
		if repo == "" {
			repo = defaultRepo
		}
		err := commands.HandlePlugin(kctlConfig, repo, pluginName, commands.ApplyOperation)
		if err != nil {
			fmt.Println(err)
			errors = true
		}
	}
	if errors {
		fmt.Printf("	Errors detected when installing plugins!\n")
	} else {
		fmt.Printf("%v	Plugins added to local deployment...\n", emoji.CheckBoxWithCheck)
	}
	time.Sleep(time.Second)
}
