package plugins

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func init() {
	// Ensure that running tests does not interfere with live system config
	cacheDirName = ".k3ai-test/cache"
}
func TestUriToFilePath(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{
			"https://raw.githubusercontent.com/argoproj/argo/stable/manifests/namespace-install.yaml",
			"raw.githubusercontent.com/argoproj/argo/stable/manifests/namespace-install.yaml",
		},
		{
			"https://test/../../../abc",
			"../../abc",
		},
		{
			"https://test.com/../../../abc",
			"../../abc",
		},
		{
			"https://test.com////abc",
			"test.com/abc",
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			actual, err := uriToFilePath(test.input)
			require.NoError(t, err)
			require.Equal(t, test.expected, actual)
		})
	}

}

func TestUriToFullPath(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
		err      error
	}{
		{
			"https://raw.githubusercontent.com/argoproj/argo/stable/manifests/namespace-install.yaml",
			"raw.githubusercontent.com/argoproj/argo/stable/manifests/namespace-install.yaml",
			nil,
		},
		{
			"https://test/../../../abc",
			"",
			errDirTraversalNotAllowed,
		},
		{
			"https://test.com/../../../abc",
			"",
			errDirTraversalNotAllowed,
		},
		{
			"https://test.com////abc",
			"test.com/abc",
			nil,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			actual, err := uriToFullPath(test.input)
			require.Equal(t, test.err, err)
			require.True(t, strings.HasSuffix(actual, test.expected))
		})
	}

}
