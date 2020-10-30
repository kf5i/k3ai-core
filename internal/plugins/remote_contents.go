package plugins

import (
	"encoding/json"
	"github.com/pkg/errors"
)

const (
	// DefaultPluginURI is the location of the plugins repository if not other location is specified
	DefaultPluginURI = "https://api.github.com/repos/kf5i/k3ai-plugins/contents/core/plugins/"
	// DefaultPluginsGroupURI is the location of the group repository if not other location is specified
	DefaultPluginsGroupURI = "https://api.github.com/repos/kf5i/k3ai-plugins/contents/core/groups/"

	dirType  = "dir"
	fileType = "file"
)

type githubContent struct {
	Name        string `json:"name"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
}

// GithubContents represents a collection of api responses from Github
type GithubContents []githubContent

func (content GithubContents) filter(filterType string) GithubContents {
	var pList GithubContents
	for _, c := range content {
		if c.Type == filterType {
			pList = append(pList, c)
		}
	}
	return pList
}

func getRepoContent(uri string) (GithubContents, error) {
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

// ContentList returns the collection of plugins in the repository
func ContentList(uri string) (GithubContents, error) {
	gHubContents, err := getRepoContent(uri)
	if err != nil {
		return nil, err
	}
	return gHubContents.filter(dirType), nil
}
