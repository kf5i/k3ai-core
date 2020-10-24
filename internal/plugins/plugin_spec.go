package plugins

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type PluginSpec struct {
	Labels     []string `yaml:",flow"`
	PluginName string   `yaml:"plugin-name"`
	Yaml       []string `yaml:"yaml"`
	Bash       []string `yaml:"bash"`
	Helm       []string `yaml:"helm"`
}

type PluginSpecs = []PluginSpec

// Encode fetches the PluginSpec
func Encode(pluginUri string) (*PluginSpec, error) {
	remoteContent, err := fetchRemoteContent(pluginUri)
	if err != nil {
		return nil, errors.Wrap(err, "error fetching plugin spec")
	}
	return unmarshal(remoteContent)
}

// validate checks for any errors in the PluginSpec
func (ps *PluginSpec) validate() error {
	return nil
}

func unmarshal(in []byte) (*PluginSpec, error) {
	var ps PluginSpec
	err := yaml.Unmarshal(in, &ps)
	if err != nil {
		return nil, err
	}
	return &ps, nil
}
