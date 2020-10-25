package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

const (
	// DefaultPluginUri is the location of the plugins repository if not other location is specified
	DefaultPluginUri = "https://api.github.com/repos/kf5i/k3ai-plugins/contents/v2/"
	dirType          = "dir"
	fileType         = "file"
)

type GithubContent struct {
	Name        string `json:"name"`
	DownloadUrl string `json:"download_url"`
	Type        string `json:"type"`
}

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

func getPluginRepoContent(uri string) (GithubContents, error) {
	const wrapMessage = "cannot load plugins"
	if uri == "" {
		uri = DefaultPluginUri
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

func GetPluginList(uri string) (GithubContents, error) {
	githubContents, err := getPluginRepoContent(uri)
	if err != nil {
		return nil, err
	}
	return githubContents.filter(dirType), nil
}

func GetPluginYamls(uri, pluginName string) (PluginSpecs, error) {
	githubContents, err := getPluginRepoContent(uri + pluginName)
	if err != nil {
		return nil, err
	}
	githubContents = githubContents.filter(fileType)
	var pluginSpecs PluginSpecs
	for _, githubContent := range githubContents {
		// Only look at the spec yaml files
		log.Printf("githubContent.Name %s", githubContent.Name )
		if strings.EqualFold(githubContent.Name, DefaultPluginFileName) {
			p, err := Encode(githubContent.DownloadUrl)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("error encoding %q", githubContent.Name))
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
