package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// These values are overridden at build time.
var (
	version   = "dev"
	commit    = "dev"
	goVersion = "unknown"
)
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: fmt.Sprintf("Print CLI version"),
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func printVersion() {
	fmt.Printf("%s version: %s\ngo version: %s\ncommit: %s\n",
		k3aiBinaryName,
		version,
		goVersion,
		commit,
	)
}
