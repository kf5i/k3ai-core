package plugins

import (
	"github.com/kf5i/k3ai-core/internal/shared"
	"github.com/pkg/errors"
)

const (
	commandFile = "file"
	// CommandKustomize is the kustomize command
	CommandKustomize = "kustomize"
	// CommandHelm is the helm command
	CommandHelm = "helm"
	// CommandDocker is the container command
	CommandDocker = "container"
	// DefaultPluginFileName is the default plugin name
	// each plugin must contain this file else it will be ignored
	DefaultPluginFileName = "plugin.yaml"
	// DefaultGroupFileName is the default group name
	// each group must contain this file else it will be ignored
	DefaultGroupFileName = "group.yaml"
)

// PluginPreType is the specification for YamlType segment of the Plugin
type PluginPreType struct {
	URL  string `yaml:"url"`
	RepoUrl string `yaml:"repourl,omitempty"`
	Type string `yaml:"type,omitempty"`
}


// YamlType is the specification for YamlType segment of the Plugin
type YamlType struct {
	URL  string `yaml:"url"`
	RepoUrl string `yaml:"repourl,omitempty"`
	Type string `yaml:"type,omitempty"`

}

//PostInstall to execute after the scripts
type PostInstall struct {
	Command string `yaml:"command,omitempty"`
}

//Plugin is the specification of each k3ai plugin
type Plugin struct {
	PluginName        string      		`yaml:"plugin-name"`
	Labels            []string    		`yaml:",flow"`
	Namespace         string      		`yaml:"namespace,omitempty"`
	PluginDescription string      		`yaml:"plugin-description"`
	PluginPreReq	  []PluginPreType   `yaml:"plugin-prerequisites,flow"`
	Yaml              []YamlType  		`yaml:"plugin-details,flow"`
	Bash              []string    		`yaml:"bash,flow"`
	Helm              []string    		`yaml:"helm,flow"`
	Container         []string    		`yaml:"container,flow"`
	PostInstall       PostInstall 		`yaml:"post-install"`
}

// Plugins list of plugins
type Plugins struct {
	Items []Plugin `yaml:"items"`
}

// Encode fetches the Plugin
func (ps *Plugin) Encode(URL string) error {
	err := encode(URL, ps)
	if err != nil {
		return err
	}
	mergeWithDefault(ps)
	return nil
}

// List fetch the plugin list
func (pls *Plugins) List(URL string) error {
	gHubContents, err := GithubContentList(URL)
	if err != nil {
		return err
	}
	for _, gHubContent := range gHubContents {
		var ps Plugin
		err := ps.Encode(shared.NormalizeURL(URL, gHubContent.Name) + DefaultPluginFileName)
		if err != nil {
			return err
		}
		pls.Items = append(pls.Items, ps)
	}
	return nil
}

// validate checks for any errors in the Plugin
func (ps *Plugin) validate() error {
	if ps.Namespace == "" {
		return errors.New("namespace value must be 'default' or another value")
	}
	for _, yamlType := range ps.Yaml {
		// First let's check if we are in the need of using Kustomize or File
		if yamlType.Type != CommandKustomize && yamlType.Type != commandFile {
			if yamlType.Type != CommandHelm && yamlType.Type != CommandDocker {
				return errors.New("type must be one of the supported ones. Please check the documentation")
			}
		}
	}
	return nil
}

func mergeWithDefault(ps *Plugin) {
	ps.Namespace = shared.GetDefaultIfEmpty(ps.Namespace, "default")
	for i, yamlTypeItem := range ps.Yaml {
		yamlType := shared.GetDefaultIfEmpty(yamlTypeItem.Type, "file")
		ps.Yaml[i] = YamlType{Type: yamlType, URL: yamlTypeItem.URL, RepoUrl: yamlTypeItem.RepoUrl}
	}
}
