package plugins

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
)

const GithubRaws = "https://raw.githubusercontent.com/kf5i/k3ai-plugins/main/"

type PluginSpec struct {
	Files      []string `yaml:",flow"`
	Labels     []string `yaml:",flow"`
	PluginName string   `yaml:"plugin-name"`
	PreInstall string   `yaml:"pre-install"`
}

func (ps *PluginSpec) Encode(uri string) (*PluginSpec, error) {

	resp, err := http.Get(GithubRaws + uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(body, ps)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return ps, nil

}
