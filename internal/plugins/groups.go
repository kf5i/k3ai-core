package plugins

import (
	"github.com/kf5i/k3ai-core/internal/shared"
	"github.com/pkg/errors"
)

//PluginGroup is the specification of a single plugin-group Item
type PluginGroup struct {
	Name    string `yaml:"name"`
	Enabled bool   `yaml:"enabled,omitempty"`
}

//Group is the specification of each k3ai plugins group
type Group struct {
	PluginType       string        `yaml:"plugin-type"`
	GroupName        string        `yaml:"group-name"`
	GroupDescription string        `yaml:"group-description"`
	Plugins          []PluginGroup `yaml:"plugins,flow"`
	InlinePlugins    []Plugin      `yaml:"inline-plugins,flow"`
}

// Groups list of Group
type Groups struct {
	Items []Group `yaml:"items"`
}

// Encode fetches the Group
func (gs *Group) Encode(URL string) error {
	return encode(URL, gs)
}

// List fetch the plugin list
func (gls *Groups) List(URL string) error {
	gHubContents, err := GithubContentList(URL)
	if err != nil {
		return err
	}
	for _, gHubContent := range gHubContents {
		var gr Group
		err := gr.Encode(shared.NormalizeURL(URL, gHubContent.Name) + DefaultGroupFileName)
		if err != nil {
			return err
		}
		gls.Items = append(gls.Items, gr)
	}
	return nil
}

func (gs *Group) validate() error {
	if len(gs.Plugins) <= 0 {
		return errors.New("empty plugins, at least on one plugin is needed")
	}
	return nil
}
