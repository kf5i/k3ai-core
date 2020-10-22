package plugins

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

const PluginUrl = "https://api.github.com/repos/kf5i/k3ai-plugins/contents/v2/"
const DirType = "dir"
const FileType = "file"

type GithubContent struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
}

type GithubContents = []GithubContent

func GetPluginsRaw(uri string) (GithubContents, error) {
	resp, err := http.Get(PluginUrl + uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	var cgs GithubContents

	err = json.Unmarshal(body, &cgs)
	if err != nil {
		errors.Wrap(err, "Can't load plugins")
		return nil, err
	}
	return cgs, nil
}

func GetPluginsFiltered(uri string, filterType string) (*GithubContents, error) {
	githubContents, err := GetPluginsRaw(uri)
	var pList GithubContents
	for _, githubContent := range githubContents {
		if githubContent.Type == filterType {
			pList = append(pList, githubContent)
		}
	}
	if err != nil {
		errors.Wrap(err, "Can't load plugins")
		return nil, err
	}
	return &pList, nil
}

func GetPluginList() (*GithubContents, error) {
	return GetPluginsFiltered("", DirType)
}

func GetPluginYamls(pluginName string) (*PluginSpecs, error) {
	var yList PluginSpecs

    githubContents, _ := GetPluginsFiltered(pluginName, FileType)
	for _, githubContent := range *githubContents {
		var pluginSpec PluginSpec
		pluginSpec.Encode(githubContent.Path)
		yList = append(yList, pluginSpec)
	}
	return &yList, nil
}
