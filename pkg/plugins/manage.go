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

//func GetPlugins() {
//	client := github.NewClient(nil)
//	repo, _, _ := client.Repositories.Get(context.Background(), "kf5i", "k3ai-plugins")
//	fmt.Printf("%s\n", *repo.ContentsURL)
//}

type GithubContent struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
	Url  string `json:"url"`
	DownloadUrl  string `json:"download_url"`
}

type Plugin struct {
	Name string `json:"name"`
	Url  string `json:"name"`
}

type Plugins struct {
	List []Plugin
}

type YamlFile struct {
	Body string
}

type YamlFiles struct {
	List []YamlFile
}

type GithubContents = []GithubContent

func GetPluginsRaw(url string, uri string) (GithubContents, error) {
	resp, err := http.Get(url + uri)
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
	githubContents, err := GetPluginsRaw(PluginUrl, uri)
	var pList Plugins
	for _, githubContent := range githubContents {
		var p Plugin
		p.Name = githubContent.Name
		p.Url = githubContent.Url
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

func GetYamls(pluginName string, pluginPath string) (*YamlFiles, error) {
	var yList YamlFiles
	githubYamls, err := GetPluginsRaw(pluginPath, "")
	for _, githubYaml := range githubYamls {
		var y YamlFile
		resp, err := http.Get(githubYaml.DownloadUrl)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			errors.Wrap(err, "Can't load yaml")
			return nil, err
		}
		y.Body = string(body)
		yList.List = append(yList.List, y)
	}
	if err != nil {
		errors.Wrap(err, "Can't load yaml")
		return nil, err
	}
	return &yList, nil
}

func GetPluginList() (*Plugins, error) {
	return GetPluginsFiltered("", DirType)
}

func GetPluginYamls(pluginName string, pluginPath string) (*YamlFiles, error) {
	return GetYamls(pluginName, pluginPath)
}
