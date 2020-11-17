package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
	"log"
	"os"
	"github.com/manifoldco/promptui"
	"strings"
	"github.com/kf5i/k3ai-core/internal/k8s/infra"
)

type pepper struct {
	Name     string
	HeatUnit int
	Peppers  int
}

func newInitCommand() *cobra.Command {
	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize K3ai Client",
		Long:  `Initialize K3ai Client, allowing user to deploy a new K8's cluster, list plugins and groups`,
		Run: func(cmd *cobra.Command, args []string) {
			/* First we check the OS to setup the env variable and than we check if everything is ready,
			 if not we do suggest the user the next steps */
			if runtime.GOOS == "windows" {
				fmt.Println("You are running on Windows")
				osFlavor := runtime.GOOS
				CheckClusterReadiness(osFlavor)
			} else if runtime.GOOS == "linux" {
				fmt.Println("You are running on Linux")
				osFlavor := runtime.GOOS
				CheckClusterReadiness(osFlavor)
			} else {
				fmt.Println("You are running on an OS other than Windows or Linux")
			}
		},
	}
	return initCmd
}

func CheckClusterReadiness(osFlavor string) {
	kubeconfig, err := os.LookupEnv("KUBECONFIG")
	if err != true {
			installK8sForMe(osFlavor)
			log.Fatal(err)
	} else {
		fmt.Println(kubeconfig)

	}
}

func installK8sForMe(osFlavor string) {
	prompt := promptui.Select{
		Label: "Select Cluster to install [default k3s]:",
		Items: []string{"K3s", "Kind", "K0s","exit"},
	}

	_, result, err := prompt.Run()
	result = strings.ToLower(result)
	
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	
	switch result {
		case "k3s":
			infra.InfraK3s(osFlavor,result)
		case "kind":
			infra.InfraKind(osFlavor,result)
		case "k0s":
			fmt.Println("k0s")
		case "exit":
			fmt.Println("okay let's exit")
		default:
			fmt.Println("okay dude let's go with k3s")

	}
}