package cli

import (

	//"os"
	// "runtime"
	// "strings"
	// "time"

	// "github.com/enescakir/emoji"
	"github.com/kf5i/k3ai-core/internal/infra/cloud"
	"github.com/kf5i/k3ai-core/internal/infra/local"

	"fmt"

	"github.com/kf5i/k3ai-core/internal/shared"
	// "github.com/manifoldco/promptui"
	//"fmt"

	"github.com/spf13/cobra"
)

const absPath = "$HOME/.k3ai/config.yaml"

var kubeconfig []byte

/* First step is to check if inside .k3ai folder exist a copy of config.yaml if not pull one from github
// the default one disabled locally with no plugins and will ask user what want to do.
// based on user choices will enable the right configuration on the config
// finally will instruct the user on how to change the config */

func newInitCommand() *cobra.Command {

	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize K3ai Client",
		Long:  `Initialize K3ai Client, allowing user to deploy a new K8's cluster, list plugins and groups`,
		Example: `k3ai init					#Will use config from $HOME/.k3ai/config.yaml and use interactive menus
k3ai init --config /myfolder/myconfig.yaml	#Use a custom config.yaml in another location(local or remote)
k3ai init --local k3s		 	#Use config target marked local and of type k3s
k3ai init --cloud civo			#Use config target marked as cloud and of type civo`,
		SilenceUsage: true,
	}

	initCmd.Flags().String("local", "", "Options availabe k3s,k0s,kind")
	initCmd.Flags().String("cloud", "", "Options availabe for cloud providers")
	initCmd.Flags().String("config", "/.k3ai/config.yaml", "Custom config file [default is $HOME/.k3ai/config.yaml]")

	initCmd.RunE = func(cmd *cobra.Command, args []string) error {
		localConfig, _ := initCmd.Flags().GetString("config")
		//localClusterConfig, _ := initCmd.Flags().GetString("local")
		remoteClusterConfig, _ := initCmd.Flags().GetString("cloud")
		if remoteClusterConfig != "" {
			cloud.CivoCloudInit("windows", "civo")
		}

		// initialize the CLI config for executing kubectl
		kctlConfig := newConfig(cmd)

		//check if config.yaml exist otherwise grab a copy
		cfg, _ := shared.Init(localConfig)
		enabled := false
		for i := range cfg.TargetCustomization {
			if cfg.TargetCustomization[i].Enabled {
				//check type call relative  function: prepare the data we need and push to the relative function
				if cfg.TargetCustomization[i].ClusterDeployment == "cloud" {
					cloud.Init(cfg.TargetCustomization[i])
				} else {
					local.Init(kctlConfig, repo, cfg.TargetCustomization[i])
				}
				enabled = true
			}
		}
		// we assume everything is false (first time?) so we need a simple interactive menu
		// meanwhile, let's warn user about the situation
		if !enabled {
			fmt.Println("No infrastructure type is marked as enabled; check your config file and try again.")
		}
		return nil
	}
	return initCmd
}
