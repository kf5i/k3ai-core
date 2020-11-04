package plugins

import (
	"fmt"
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

//Plugin is the specification of each k3ai plugin
type Plugin struct {
	Namespace  string     `yaml:"namespace,omitempty"`
	Labels     []string   `yaml:",flow"`
	PluginName string     `yaml:"plugin-name"`
	Yaml       []YamlType `yaml:"yaml,flow"`
	Bash       []string   `yaml:"bash,flow"`
	Helm       []string   `yaml:"helm,flow"`
}

// Plugins is a Plugin collection
type Plugins struct {
	Plugins []Plugin
}

// Encode fetches the Plugin
func (Plugin) Encode(pluginURI string) (*Plugin, error) {
	remoteContent, err := FetchFromSourceURI(pluginURI)
	if err != nil {
		return nil, errors.Wrap(err, "error fetching plugins")
	}
	var ps Plugin
	err = yaml.Unmarshal(remoteContent, &ps)
	mergeWithDefault(&ps)
	if err != nil {
		return nil, err
	}
	return &ps, nil
}

// Encode fetches the Plugins
func (Plugins) Encode(pluginURI string, pluginName string) (*Plugins, error) {
	var plugins Plugins
	if !isHTTP(pluginURI) {
		var p Plugin
		r, err := p.Encode(shared.NormalizePath(DefaultPluginFileName, pluginURI, pluginName))
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error encoding %q", pluginURI))
		}

		plugins.Plugins = append(plugins.Plugins, *r)
		return &plugins, nil
	}

	gHubContents, err := getRepoContent(getDefaultIfEmpty(pluginURI, DefaultPluginURI) + pluginName)
	if err != nil {
		return nil, err
	}
	gHubContents = gHubContents.filter(fileType)
	for _, githubContent := range gHubContents {
		if githubContent.Name == DefaultPluginFileName {
			var p Plugin
			r, err := p.Encode(githubContent.DownloadURL)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("error encoding %q", githubContent.Name))
			}
			plugins.Plugins = append(plugins.Plugins, *r)
		}
	}

	return &plugins, nil
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
	ps.Namespace = getDefaultIfEmpty(ps.Namespace, "default")
	for i, yamlTypeItem := range ps.Yaml {
		yamlType := getDefaultIfEmpty(yamlTypeItem.Type, "file")
		ps.Yaml[i] = YamlType{Type: yamlType, URL: yamlTypeItem.URL}
	}
}
