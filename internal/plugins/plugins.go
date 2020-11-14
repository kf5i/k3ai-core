package plugins

import (
	"github.com/kf5i/k3ai-core/internal/shared"
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
	// DefaultGroupFileName is the default group name
	// each group must contain this file else it will be ignored
	DefaultGroupFileName = "group.yaml"
)

// YamlType is the specification for YamlType segment of the Plugin
type YamlType struct {
	URL  string `yaml:"url"`
	Type string `yaml:"type,omitempty"`
}

//PostInstall to execute after the scripts
type PostInstall struct {
	Command string `yaml:"command,omitempty"`
}

//Plugin is the specification of each k3ai plugin
type Plugin struct {
	Namespace   string      `yaml:"namespace,omitempty"`
	Labels      []string    `yaml:",flow"`
	PluginName  string      `yaml:"plugin-name"`
	Yaml        []YamlType  `yaml:"yaml,flow"`
	Bash        []string    `yaml:"bash,flow"`
	Helm        []string    `yaml:"helm,flow"`
	PostInstall PostInstall `yaml:"post-install"`
}

// Encode fetches the Plugin
func (ps *Plugin) Encode(pluginURI string) error {

	if !isHTTP(pluginURI) {
		remoteContent, err := FetchFromSourceURI(pluginURI)
		err = yaml.Unmarshal(remoteContent, &ps)
		mergeWithDefault(ps)
		if err != nil {
			return err
		}
		return nil
	}

	gHubContent, err := getRepoContent(pluginURI)
	if err != nil {
		return errors.Wrap(err, "error fetching plugins content")
	}
	remoteContent, err := FetchFromSourceURI(gHubContent.DownloadURL)
	if err != nil {
		return errors.Wrap(err, "error fetching plugins")
	}
	err = yaml.Unmarshal(remoteContent, &ps)
	mergeWithDefault(ps)
	if err != nil {
		return err
	}
	return nil
}

// validate checks for any errors in the Plugin
func (ps *Plugin) validate() error {
	if ps.Namespace == "" {
		return errors.New("namespace value must be 'default' or another value")
	}
	for _, yamlType := range ps.Yaml {
		if yamlType.Type != CommandKustomize && yamlType.Type != commandFile {
			return errors.New("type must be file or kustomize")
		}
	}
	return nil
}

func mergeWithDefault(ps *Plugin) {
	ps.Namespace = shared.GetDefaultIfEmpty(ps.Namespace, "default")
	for i, yamlTypeItem := range ps.Yaml {
		yamlType := shared.GetDefaultIfEmpty(yamlTypeItem.Type, "file")
		ps.Yaml[i] = YamlType{Type: yamlType, URL: yamlTypeItem.URL}
	}
}
