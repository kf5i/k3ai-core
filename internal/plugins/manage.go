package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

const (
	// DefaultPluginURI is the location of the plugins repository if not other location is specified
	DefaultPluginURI = "https://api.github.com/repos/kf5i/k3ai-plugins/contents/v2/"
	dirType          = "dir"
	fileType         = "file"
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

func getPluginRepoContent(uri string) (GithubContents, error) {
	const wrapMessage = "cannot load plugins"
	if uri == "" {
		uri = DefaultPluginURI
	}
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

// PluginList returns the collection of plugins in the repository
func PluginList(uri string) (GithubContents, error) {
	gHubContents, err := getPluginRepoContent(uri)
	if err != nil {
		return nil, err
	}
	return gHubContents.filter(dirType), nil
}

// PluginYamls inflates the PluginSpecs for a plugin
func PluginYamls(uri, pluginName string) (PluginSpecs, error) {
	gHubContents, err := getPluginRepoContent(uri + pluginName)
	if err != nil {
		return nil, err
	}
	gHubContents = gHubContents.filter(fileType)
	var pluginSpecs PluginSpecs
	for _, gHubContent := range gHubContents {
		// Only look at the spec yaml files
		if strings.HasSuffix(gHubContent.Name, ".yaml") {
			p, err := Encode(gHubContent.DownloadURL)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("error encoding %q", gHubContent.Name))
			}
			pluginSpecs = append(pluginSpecs, *p)
		}
	}
	return pluginSpecs, nil
}

func fetchRemoteContent(uri string) ([]byte, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	// TODO: Check http status code for better error messages
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
