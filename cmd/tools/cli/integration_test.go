// +build integration

package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func setUp() (*cobra.Command, *bytes.Buffer) {
	cmd := &cobra.Command{
		Use:           "k3ai-cli-test",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	out := bytes.NewBuffer(nil)
	cmd.SetOut(out)
	cmd.SetErr(out)
	setupCli(cmd)
	return cmd, out
}
func TestApply(t *testing.T) {
	cmd, out := setUp()
	cmd.SetArgs([]string{"apply", "--kubectl", "ci-tests"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	assertMessage(t, out.String(), `service/argo-server created`)
}

func assertMessage(t *testing.T, input string, message string) {
	if !strings.Contains(input, message) {
		t.Fatalf("did not find %q in %q", message, input)
	}
}
func TestDelete(t *testing.T) {
	cmd, out := setUp()
	cmd.SetOut(out)
	cmd.SetErr(out)
	cmd.SetArgs([]string{"delete", "--kubectl", "ci-tests"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	assertMessage(t, out.String(), `service "argo-server" deleted`)
}

func TestLocalApply(t *testing.T) {
	cmd, out := setUp()
	cmd.SetArgs([]string{"apply", "argo-workflow-ns", "--kubectl", "--plugin-repo", joinWithRootData("plugins")})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	assertMessage(t, out.String(), `service/argo-server created`)
}

func TestLocalDelete(t *testing.T) {
	cmd, out := setUp()
	cmd.SetArgs([]string{"delete", "argo-workflow-ns", "--kubectl", "--plugin-repo", joinWithRootData("plugins")})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	assertMessage(t, out.String(), `service "argo-server" deleted`)
}

func TestLocalGroupApply(t *testing.T) {
	cmd, out := setUp()
	cmd.SetArgs([]string{"apply", "-g", "argo", "--kubectl", "--group-repo", joinWithRootData("groups")})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	assertMessage(t, out.String(), `service/argo-server created`)
}

func TestLocalGroupDelete(t *testing.T) {
	cmd, out := setUp()
	cmd.SetArgs([]string{"delete", "-g", "argo", "--kubectl", "--group-repo", joinWithRootData("groups")})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	assertMessage(t, out.String(), `service "argo-server" deleted`)
}

func getRootTestData() string {
	return "../../../local_repo/core/"
}

func joinWithRootData(fileURI string) string {
	return getRootTestData() + fileURI
}
