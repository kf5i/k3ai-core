package plugins

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//PluginGroupSpec is the specification of a single plugin-group Item
type PluginGroupSpec struct {
	Name    string `yaml:"name"`
	Enabled bool   `yaml:"enabled,omitempty"`
}

//PluginsGroupSpec is the specification of each k3ai plugins group
type PluginsGroupSpec struct {
	PluginType string            `yaml:"plugin-type"`
	GroupName  string            `yaml:"group-name"`
	Plugins    []PluginGroupSpec `yaml:"plugins,flow"`
}

// EncodePluginsGroupSpec fetches the PluginsGroupSpec
func EncodePluginsGroupSpec(pluginsGroupURI string) (*PluginsGroupSpec, error) {
	remoteContent, err := fetchRemoteContent(pluginsGroupURI)
	if err != nil {
		return nil, errors.Wrap(err, "error fetching plugins group spec")
	}

	return unmarshalPluginsGroupSpec(remoteContent)
}

func (ps *PluginsGroupSpec) validatePluginsGroupSpec() error {
	return nil
}

func LoadPluginsGroupSpecFormFile(filePath string) (*PluginsGroupSpec, error) {
	yamlContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return unmarshalPluginsGroupSpec(yamlContent)
}

func unmarshalPluginsGroupSpec(in []byte) (*PluginsGroupSpec, error) {
	var ps PluginsGroupSpec
	err := yaml.Unmarshal(in, &ps)

	if err != nil {
		return nil, err
	}
	return &ps, nil
}
