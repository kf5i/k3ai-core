package plugins

import (
	"fmt"
	"github.com/kf5i/k3ai-core/internal/shared"
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

//Groups is the specification of each k3ai plugins group
type Groups struct {
	Groups []Group
}

// Encode fetches the Group
func (gs Group) Encode(groupURI string) (*Group, error) {
	remoteContent, err := FetchFromSourceURI(groupURI)
	if err != nil {
		return nil, errors.Wrap(err, "error fetching plugins group spec")
	}
	err = yaml.Unmarshal(remoteContent, &gs)
	if err != nil {
		return nil, err
	}
	return &gs, nil
}

func (gs *Group) validate() error {
	if len(gs.Plugins) <= 0 {
		return errors.New("empty plugins, at least on one plugin is needed")
	}

	return nil
}

// Encode fetches the Plugins
func (groups Groups) Encode(groupURI string, groupName string) (*Groups, error) {
	if !isHTTP(groupURI) {
		var p Group
		r, err := p.Encode(shared.NormalizePath(DefaultGroupFileName, groupURI, groupName))
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error encoding %q", groupURI))
		}

		groups.Groups = append(groups.Groups, *r)
		return &groups, nil
	}

	gHubContents, err := getRepoContent(shared.GetDefaultIfEmpty(groupURI, DefaultPluginsGroupURI) + groupName)
	if err != nil {
		return nil, err
	}
	gHubContents = gHubContents.filter(fileType)
	for _, githubContent := range gHubContents {
		if githubContent.Name == DefaultGroupFileName {
			var p Group
			r, err := p.Encode(githubContent.DownloadURL)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("error encoding %q", githubContent.Name))
			}
			groups.Groups = append(groups.Groups, *r)
		}
	}

	return &groups, nil
}
