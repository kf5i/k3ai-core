package cli

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/kf5i/k3ai-core/internal/k8s/kctl"
	"github.com/kf5i/k3ai-core/internal/plugins"
	"github.com/spf13/cobra"
)

const k3aiBinaryName = "k3ai-cli"

var rootCmd = &cobra.Command{
	Use:   k3aiBinaryName,
	Short: fmt.Sprintf(`%s installs AI tools`, k3aiBinaryName),
	Long: fmt.Sprintf(` %s is a lightweight infrastructure-in-a-box solution specifically built to
	install and configure AI tools and platforms in production environments on Edge
	and IoT devices as easily as local test environments.`, k3aiBinaryName),
	SilenceUsage:  true,
	SilenceErrors: true,
}

var (
	pluginRepoURI       string
	pluginsGroupRepoURI string
	useKubectl          bool
)

func init() {
	setupCli(rootCmd)
}

func setupCli(baseCmd *cobra.Command) {
	baseCmd.PersistentFlags().StringVarP(&pluginRepoURI, "plugin-repo", "", plugins.DefaultPluginURI, "URI for the plugins repository. Must begin with https:// or file://")
	baseCmd.PersistentFlags().StringVarP(&pluginsGroupRepoURI, "group-repo", "", plugins.DefaultPluginsGroupURI, "URI for the plugins repository. Must begin with https:// or file://")
	baseCmd.PersistentFlags().BoolVarP(&useKubectl, "kubectl", "", false, "Use kubectl for deployment. Uses k3s when set to false")
	baseCmd.AddCommand(versionCmd)
	baseCmd.AddCommand(newApplyCommand())

	baseCmd.AddCommand(newDeleteCommand())
	baseCmd.AddCommand(newListCommand())
}

//Execute is the entrypoint of the commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type config struct {
	context.Context
	stdin      io.Reader // standard input
	stdout     io.Writer // standard output
	stderr     io.Writer // standard error
	useKubectl bool
}

func newConfig(cmd *cobra.Command) kctl.Config {
	return &config{
		context.Background(),
		cmd.InOrStdin(), cmd.OutOrStdout(), cmd.ErrOrStderr(),
		useKubectl,
	}
}

func (c *config) Stdin() io.Reader {
	return c.stdin
}
func (c *config) Stdout() io.Writer {
	return c.stdout
}
func (c *config) Stderr() io.Writer {
	return c.stderr
}

func (c *config) UseKubectl() bool {
	return c.useKubectl
}
