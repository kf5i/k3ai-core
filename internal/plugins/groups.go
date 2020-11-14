package plugins

import (
	"github.com/pkg/errors"
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
func (gs *Group) Encode(URL string) error {
	return encode(URL, gs)
}

func (gs *Group) validate() error {
	if len(gs.Plugins) <= 0 {
		return errors.New("empty plugins, at least on one plugin is needed")
	}

	return nil
}
