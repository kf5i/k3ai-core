package plugins

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
)

type PluginSpec struct {
	Files      []string `yaml:",flow"`
	Labels     []string `yaml:",flow"`
	PluginName string   `yaml:"plugin-name"`
	PreInstall string   `yaml:"pre-install"`
}

func (ps *PluginSpec) Encode(url string) (*PluginSpec, error) {

	resp, err := http.Get(url)
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
