package plugins

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
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
	remoteContent, err := fetchRemoteContent(PluginUrl + uri)
	if err != nil {
	}
	var cgs GithubContents

	err = json.Unmarshal(remoteContent, &cgs)
	if err != nil {
		errors.Wrap(err, "cannot load plugins")
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
		errors.Wrap(err, "cannot load plugins")
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
		p, err := Encode(githubContent.Path)
		if err != nil {
			return nil, err
		}
		yList = append(yList, *p)
	}
	return &yList, nil
}

func fetchRemoteContent(uri string) ([]byte, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
