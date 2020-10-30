// +build integration

package cli

import (
	"bytes"
	"log"
	"testing"

	"github.com/spf13/cobra"
)

var cmd *cobra.Command

func setUp() {
	cmd = &cobra.Command{
		Use:           "k3ai-cli-test",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	setupCli(cmd)
}
func TestApply(t *testing.T) {
	setUp()
	out := bytes.NewBuffer(nil)
	cmd.SetOut(out)
	cmd.SetErr(out)
	cmd.SetArgs([]string{"apply", "--kubectl", "argo"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	log.Println(out.String())
}

func TestDelete(t *testing.T) {
	setUp()
	out := bytes.NewBuffer(nil)
	cmd.SetOut(out)
	cmd.SetErr(out)
	cmd.SetArgs([]string{"delete", "--kubectl", "argo"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	log.Println(out.String())
}
