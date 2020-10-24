package cli

import (
	"context"
	"fmt"
	"io"
	"os"

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
}

func init() {

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(applyCmd)
	rootCmd.AddCommand(listCmd)
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
	stdin  io.Reader // standard input
	stdout io.Writer // standard output
	stderr io.Writer // standard error
}

func newConfig() kctl.Config {
	return &config{
		context.Background(),
		os.Stdin, os.Stdout, os.Stderr,
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
