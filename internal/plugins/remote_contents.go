package plugins

import (
	"encoding/json"
	"github.com/pkg/errors"
)

const (
	// DefaultRepo is the location of the plugins repository if not other location is specified
	DefaultRepo = "https://api.github.com/repos/kf5i/k3ai-plugins/contents/core/"

	PluginDir = "plugins"
	GroupsDir = "groups"

	DefaultPluginsGroupURI = DefaultRepo + GroupsDir
	DefaultPluginsURI      = DefaultRepo + PluginDir

	dirType  = "dir"
	fileType = "file"
)

type GithubContent struct {
	Name        string `json:"name"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
}

// GithubContents represents a collection of api responses from Github
type GithubContents []GithubContent

func (content GithubContents) filter(filterType string) GithubContents {
	var pList GithubContents
	for _, c := range content {
		if c.Type == filterType {
			pList = append(pList, c)
		}
	}
	return pList
}

func getRepoContents(uri string) (GithubContents, error) {
	const wrapMessage = "cannot load plugins"

	remoteContent, err := fetchRemoteContent(uri)
	if err != nil {
		return nil, errors.Wrap(err, wrapMessage)
	}
	var cgs GithubContents
	err = json.Unmarshal(remoteContent, &cgs)
	if err != nil {
		return nil, errors.Wrap(err, wrapMessage)
	}
	return cgs, nil
}

func getRepoContent(uri string) (*GithubContent, error) {
	const wrapMessage = "cannot load plugins"

	remoteContent, err := fetchRemoteContent(uri)
	if err != nil {
		return nil, errors.Wrap(err, wrapMessage)
	}
	var cgs GithubContent
	err = json.Unmarshal(remoteContent, &cgs)
	if err != nil {
		return nil, errors.Wrap(err, wrapMessage)
	}
	return &cgs, nil
}

// ContentList returns the collection of plugins in the repository
func ContentList(uri string) (GithubContents, error) {
	gHubContents, err := getRepoContents(uri)
	if err != nil {
		return nil, err
	}
	return gHubContents.filter(dirType), nil
}
