package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const k3aiBinaryName = "k3ai-cli"

var rootCmd = &cobra.Command{
	Use:   k3aiBinaryName,
	Short: fmt.Sprintf(`%s installs AI tools`, k3aiBinaryName),
	Long: fmt.Sprintf(` %s is a lightweight infrastructure-in-a-box solution specifically built to
	install and configure AI tools and platforms in production environments on Edge
	and IoT devices as easily as local test environments.`, k3aiBinaryName),
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(applyCmd)
}

//Execute is the entrypoint of the commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
