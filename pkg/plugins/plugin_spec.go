package plugins

import (
	"log"

	"gopkg.in/yaml.v2"
)

const GithubRaws = "https://raw.githubusercontent.com/kf5i/k3ai-plugins/main/"

type PluginSpec struct {
	// Files is deprecated and would be dropped in favor of
	// specific types of yaml, bash, helm.
	Files      []string `yaml:",flow"`
	Labels     []string `yaml:",flow"`
	PluginName string   `yaml:"plugin-name"`
	PreInstall string   `yaml:"pre-install"`
	Yaml       []string `yaml:"yaml"`
	Bash       []string `yaml:"bash"`
	Helm       []string `yaml:"helm"`
}

type PluginSpecs = []PluginSpec

func Encode(pluginPath string) (*PluginSpec, error) {
	remoteContent, err := fetchRemoteContent(GithubRaws + pluginPath)
	if err != nil {
		log.Printf("error fetching plugin spec  #%v ", err)
		return nil, err
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
		log.Printf("unmarshal: %v", err)
		return nil, err
	}
	return &ps, nil
}
