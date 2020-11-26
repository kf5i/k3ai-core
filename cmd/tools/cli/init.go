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
	"time"

	"github.com/enescakir/emoji"
	"github.com/kf5i/k3ai-core/internal/infra"
	"github.com/kf5i/k3ai-core/internal/shared"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type pepper struct {
	Name     string
	HeatUnit int
	Peppers  int
}

func newInitCommand() *cobra.Command {
	//var p string
	//var c string
	var localCluster, remoteCluster bool // used for flags
	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize K3ai Client",
		Long:  `Initialize K3ai Client, allowing user to deploy a new K8's cluster, list plugins and groups`,
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.Name() == "init" && len(args) <= 0 && localCluster == false && remoteCluster == false {
				osFlavor := runtime.GOOS
				checkClusterReadiness(osFlavor)
			}

			if localCluster && remoteCluster {
				// print localCluster and build date
				fmt.Println("Not yet supported")
			} else if localCluster {
				// print only localCluster
				var a = args
				if len(args) == 0 {
					osFlavor := runtime.GOOS
					checkClusterReadiness(osFlavor)
				} else {
					for i := 0; i < len(a); i++ {
						osFlavor := runtime.GOOS
						switch a[i] {
						case "k3s":
							infra.K3s(osFlavor, a[i])
						case "k0s":
							infra.K0s(osFlavor, a[i])
						case "kind":
							infra.Kind(osFlavor, a[i])
						default:
							checkClusterReadiness(osFlavor)
						}
					}
				}
			} else if remoteCluster {
				// print only remoteCluster
				var a = args
				if len(args) <= 0 {
					osFlavor := runtime.GOOS
					installRemoteK8sForMe(osFlavor)
				} else {
					for i := 0; i < len(a); i++ {
						a[i] = strings.ToLower(a[i])
						osFlavor := runtime.GOOS
						switch a[i] {
						case "civo":
							infra.CloudProviders(osFlavor, a[i])
						case "azure":
							infra.K0s(osFlavor, a[i])
						case "google":
							infra.Kind(osFlavor, a[i])
						case "aws":
							infra.Kind(osFlavor, a[i])
						default:
							checkClusterReadiness(osFlavor)
						}
					}
				}
			}
		},
	}

	initCmd.Flags().BoolVar(&localCluster, "local", false, "Options availabe k3s,k0s,kind")
	initCmd.Flags().BoolVar(&remoteCluster, "cloud", false, "Options availabe for cloud providers")
	return initCmd
}

// checkCluserReadiness check the KUBECONFIG existance
func checkClusterReadiness(osFlavor string) {
	kubepath := "/usr/local/bin/kubectl"
	// let's first check if kubectl exist on the current flavor

	if osFlavor == "windows" {
		kubepath = "C:/Windows/System32/kubectl.exe"
	}
	kubeExist := shared.CheckKubectl(osFlavor, kubepath)
	if kubeExist == false {
		fmt.Printf("%v It seem you don't have kubectl installed", emoji.PensiveFace)
		fmt.Printf("%v Please head to: https://kubernetes.io/docs/tasks/tools/install-kubectl/ for more informations", emoji.Information)
		fmt.Printf("Thank you for using K3ai %v\n", emoji.WavingHand)
		time.Sleep(3 * time.Second)
		os.Exit(0)
	}
	// if kubectl is there let see if there's a KUBECONFIG configured
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
		Label: "Select you local cluster flavor [default k3s]:",
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
		fmt.Printf("Thank you for using K3ai %v\n", emoji.WavingHand)
		os.Exit(0)
	default:
		os.Exit(0)
	}
}

func installRemoteK8sForMe(osFlavor string) {
	prompt := promptui.Select{
		Label: "Select Remote Cluster to install [default Civo]:",
		Items: []string{"Civo", "Azure", "Google", "AWS", "exit"},
	}

	_, result, err := prompt.Run()
	result = strings.ToLower(result)

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "Civo":
		infra.CloudProviders(osFlavor, result)
	case "Azure":
		fmt.Println("Azure is not supported yet...")
		os.Exit(0)
	case "Google":
		fmt.Println("Google is not supported yet...")
		os.Exit(0)
	case "AWS":
		fmt.Println("AWS is not supported yet...")
		os.Exit(0)
	case "exit":
		fmt.Printf("Thank you for using K3ai %v\n", emoji.WavingHand)
		os.Exit(0)
	default:
		os.Exit(0)

	}
}
