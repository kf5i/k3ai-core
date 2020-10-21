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

//func GetPlugins() {
//	client := github.NewClient(nil)
//	repo, _, _ := client.Repositories.Get(context.Background(), "kf5i", "k3ai-plugins")
//	fmt.Printf("%s\n", *repo.ContentsURL)
//}

type GithubContent struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
}

type Plugin struct {
	Name string `json:"name"`
}

type Plugins struct {
	List []Plugin
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

func GetPluginsFiltered(uri string, filterType string) (*Plugins, error) {
	githubContents, err := GetPluginsRaw(uri)
	var pList Plugins
	for _, githubContent := range githubContents {
		var p Plugin
		p.Name = githubContent.Name
		if githubContent.Type == filterType {
			pList.List = append(pList.List, p)
		}
	}
	if err != nil {
		errors.Wrap(err, "Can't load plugins")
		return nil, err
	}
	return &pList, nil
}

func GetPluginList() (*Plugins, error) {
	return GetPluginsFiltered("", DirType)
}

func GetPluginYamls(pluginName string) (*Plugins, error) {
	return GetPluginsFiltered(pluginName, FileType)
}
