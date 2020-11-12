package cli

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kf5i/k3ai-core/internal/settings"

	"github.com/kf5i/k3ai-core/internal/k8s/kctl"
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
	s, err := settings.LoadSettingFormHomeFile()
	if err != nil {
		log.Fatalf("can't read settings")
	}

	baseCmd.PersistentFlags().StringVarP(&pluginRepoURI, "plugin-repo", "", s.PluginRepo, "URI for the plugins repository. ")
	baseCmd.PersistentFlags().StringVarP(&pluginsGroupRepoURI, "group-repo", "", s.GroupRepo, "URI for the plugin groups repository")
	baseCmd.PersistentFlags().BoolVarP(&useKubectl, "kubectl", "", s.UseKubectl, "Use kubectl for deployment. Uses k3s when set to false")
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
