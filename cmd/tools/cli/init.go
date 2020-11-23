package cli

/*Author: Alessandro Festa
Infra package allow a user to install a local cluster based on:
K3s
K0s
Kind
*/
import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/kf5i/k3ai-core/internal/infra"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type pepper struct {
	Name     string
	HeatUnit int
	Peppers  int
}

func newInitCommand() *cobra.Command {
	var p string
	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize K3ai Client",
		Long:  `Initialize K3ai Client, allowing user to deploy a new K8's cluster, list plugins and groups`,
		Run: func(cmd *cobra.Command, args []string) {
			/* First we check the OS to setup the env variable and than we check if everything is ready,
			if not we do suggest the user the next steps */

			p = strings.ReplaceAll(p, " ", "")
			if runtime.GOOS == "windows" {
				fmt.Println("You are running on Windows")
				osFlavor := runtime.GOOS
				if p == "" {
					checkClusterReadiness(osFlavor)
				} else {
					switch p {
					case "k3s":
						infra.K3s(osFlavor, p)
					case "k0s":
						infra.K0s(osFlavor, p)
					case "kind":
						infra.Kind(osFlavor, p)
					default:
						checkClusterReadiness(osFlavor)
					}

				}
			} else if runtime.GOOS == "linux" {
				fmt.Println("You are running on Linux")
				osFlavor := runtime.GOOS
				if p == "" {
					checkClusterReadiness(osFlavor)
				} else {
					switch p {
					case "k3s":
						infra.K3s(osFlavor, p)
					case "k0s":
						infra.K0s(osFlavor, p)
					case "kind":
						infra.Kind(osFlavor, p)
					default:
						infra.K3s(osFlavor, p)
					}
				}
			} else if runtime.GOOS == "darwin" {
				fmt.Println("You are running on MacOs")
				osFlavor := runtime.GOOS
				infra.Kind(osFlavor, p)
			} else {
				fmt.Println("You are running on an OS other than Windows or Linux")
			}
		},
	}

	initCmd.Flags().StringVarP(&p, "cluster", "c", "", "Options availabe k3s,k0s,kind")
	return initCmd
}

// checkCluserReadiness check the KUBECONFIG existance
func checkClusterReadiness(osFlavor string) {
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
		Items: []string{"K3s", "Kind", "K0s", "exit"},
	}

	_, result, err := prompt.Run()
	result = strings.ToLower(result)

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "k3s":
		infra.K3s(osFlavor, result)
	case "kind":
		infra.Kind(osFlavor, result)
	case "k0s":
		infra.K0s(osFlavor, result)
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("okay dude let's go with k3s")

	}
}
