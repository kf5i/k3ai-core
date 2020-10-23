package cmd

import (
	"fmt"
	"os"

	"github.com/kf5i/k3ai-core/pkg/k8s/kctl"
	"github.com/kf5i/k3ai-core/pkg/plugins"
	"github.com/spf13/cobra"
)

const k3aiBinaryName = "k3ai-client"

var rootCmd = &cobra.Command{
	Use:   k3aiBinaryName,
	Short: fmt.Sprintf(`%s installs AI tools`, k3aiBinaryName),
	Long: fmt.Sprintf(` %s is a lightweight infrastructure-in-a-box solution specifically built to
	install and configure AI tools and platforms in production environments on Edge
	and IoT devices as easily as local test environments.`, k3aiBinaryName),
	RunE: func(cmd *cobra.Command, args []string) error {
		pluginList, _ := plugins.GetPluginList()
		fmt.Printf("Plugin list: %s\n", pluginList)

		pluginSpecList, _ := plugins.GetPluginYamls("argo")
		for _, pluginSpec := range *pluginSpecList {
			//fmt.Printf("Files: %s, Path %s \n", githubContent.Name, githubContent.Path)
			fmt.Printf("Plugin YAML content: %s, name: %s \n", pluginSpec.Files, pluginSpec.PluginName)
			fmt.Println("Going to Apply the Apply")
			kctl.ApplyFiles(pluginSpec, nil)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

//Execute is the entrypoint of the commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
