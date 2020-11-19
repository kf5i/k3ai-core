package plugins

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TesturiToFilePath(t *testing.T) {
	var test = "https://raw.githubusercontent.com/argoproj/argo/stable/manifests/namespace-install.yaml"
	expected := "raw.githubusercontent.com/argoproj/argo/stable/manifests/namespace-install.yaml"
	actual, err := uriToFilePath(test)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}
