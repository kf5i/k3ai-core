package plugins

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

//PluginGroup is the specification of a single plugin-group Item
type PluginGroup struct {
	Name    string `yaml:"name"`
	Enabled bool   `yaml:"enabled,omitempty"`
}

//Group is the specification of each k3ai plugins group
type Group struct {
	PluginType    string        `yaml:"plugin-type"`
	GroupName     string        `yaml:"group-name"`
	Plugins       []PluginGroup `yaml:"plugins,flow"`
	InlinePlugins []Plugin      `yaml:"inline-plugins,flow"`
}

// Encode fetches the Group
func (gs *Group) Encode(groupURI string) error {
	if !isHTTP(groupURI) {
		remoteContent, err := FetchFromSourceURI(groupURI)
		err = yaml.Unmarshal(remoteContent, &gs)
		if err != nil {
			return err
		}
		return nil
	}

	gHubContent, err := getRepoContent(groupURI)
	if err != nil {
		return errors.Wrap(err, "error fetching plugins content")
	}
	remoteContent, err := FetchFromSourceURI(gHubContent.DownloadURL)
	if err != nil {
		return errors.Wrap(err, "error fetching plugins group spec")
	}
	err = yaml.Unmarshal(remoteContent, &gs)
	if err != nil {
		return err
	}
	return nil
}

func (gs *Group) validate() error {
	if len(gs.Plugins) <= 0 {
		return errors.New("empty plugins, at least on one plugin is needed")
	}

	return nil
}
