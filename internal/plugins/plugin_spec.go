package plugins

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	CommandFile           = "file"
	CommandKustomize      = "kustomize"
	DefaultPluginFileName = "plugin.yaml"
)

type YamlSpec struct {
	Url       string `yaml:"url"`
	NameSpace string `yaml:"namespace,omitempty"`
	Type      string `yaml:"type,omitempty"`
}

type PluginSpec struct {
	Labels     []string   `yaml:",flow"`
	PluginName string     `yaml:"plugin-name"`
	Yaml       []YamlSpec `yaml:"yaml"`
	Bash       []string   `yaml:"bash"`
	Helm       []string   `yaml:"helm"`
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
	for _, spec := range ps.Yaml {
		if spec.Type != CommandKustomize && spec.Type != CommandFile {
			return errors.New("type must be file or kustomize")
		}
		if spec.NameSpace == "" {
			return errors.New("namespace value must be 'default' or another value")
		}
	}
	return nil
}

func unmarshal(in []byte) (*PluginSpec, error) {
	var ps PluginSpec
	err := yaml.Unmarshal(in, &ps)
	mergeWithDefault(ps)
	if err != nil {
		return nil, err
	}
	return &ps, nil
}

func mergeWithDefault(ps PluginSpec) {
	for i, spec := range ps.Yaml {
		nameSpace := spec.NameSpace
		yamlType := spec.Type
		if spec.NameSpace == "" {
			nameSpace = "default"
		}
		if spec.Type == "" {
			yamlType = "file"
		}
		ps.Yaml[i] = YamlSpec{Type: yamlType, NameSpace: nameSpace, Url: spec.Url}
	}
}
