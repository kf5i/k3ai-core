package plugins

import (
	"fmt"
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
	PluginType string        `yaml:"plugin-type"`
	GroupName  string        `yaml:"group-name"`
	Plugins    []PluginGroup `yaml:"plugins,flow"`
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
	return nil
}

// Encode fetches the Plugins
func (Groups) Encode(groupURI string, groupName string) (*Groups, error) {
	var groups Groups
	if !isHTTP(groupURI) {
		var p Group
		r, err := p.Encode(NormalizePath(groupURI, groupName, DefaultGroupFileName))
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error encoding %q", groupURI))
		}

		groups.Groups = append(groups.Groups, *r)
		return &groups, nil
	}

	gHubContents, err := getRepoContent(setDefaultIfEmpty(groupURI, DefaultPluginsGroupURI) + groupName)
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
