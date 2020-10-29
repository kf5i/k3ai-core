package plugins

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	commandFile = "file"
	// CommandKustomize is the kustomize command
	CommandKustomize = "kustomize"
	// DefaultPluginFileName is the default plugin name
	// each plugin must contain this file else it will be ignored
	DefaultPluginFileName = "plugin.yaml"
)

// YamlSpec is the specification for Yaml segment of the PluginSpec
type YamlSpec struct {
	URL  string `yaml:"url"`
	Type string `yaml:"type,omitempty"`
}

//PluginSpec is the specification of each k3ai plugin
type PluginSpec struct {
	Namespace  string     `yaml:"namespace,omitempty"`
	Labels     []string   `yaml:",flow"`
	PluginName string     `yaml:"plugin-name"`
	Yaml       []YamlSpec `yaml:"yaml,flow"`
	Bash       []string   `yaml:"bash,flow"`
	Helm       []string   `yaml:"helm,flow"`
}

// PluginSpecs is a PluginSpec collection
type PluginSpecs = []PluginSpec

// EncodePluginSpec fetches the PluginSpec
func EncodePluginSpec(pluginURI string) (*PluginSpec, error) {
	remoteContent, err := fetchRemoteContent(pluginURI)
	if err != nil {
		return nil, errors.Wrap(err, "error fetching plugin spec")
	}

	return unmarshalPluginSpec(remoteContent)
}

// validatePluginSpec checks for any errors in the PluginSpec
func (ps *PluginSpec) validatePluginSpec() error {
	if ps.Namespace == "" {
		return errors.New("namespace value must be 'default' or another value")
	}
	for _, spec := range ps.Yaml {
		if spec.Type != CommandKustomize && spec.Type != commandFile {
			return errors.New("type must be file or kustomize")
		}
	}
	return nil
}

func unmarshalPluginSpec(in []byte) (*PluginSpec, error) {
	var ps PluginSpec
	err := yaml.Unmarshal(in, &ps)
	mergeWithDefault(&ps)

	if err != nil {
		return nil, err
	}
	return &ps, nil
}

func mergeWithDefault(ps *PluginSpec) {
	if ps.Namespace == "" {
		ps.Namespace = "default"
	}
	for i, spec := range ps.Yaml {
		yamlType := spec.Type
		if spec.Type == "" {
			yamlType = "file"
		}
		ps.Yaml[i] = YamlSpec{Type: yamlType, URL: spec.URL}
	}
}
