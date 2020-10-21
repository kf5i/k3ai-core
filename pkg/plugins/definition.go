package plugins

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type PluginSpec struct {
	Files      []string `yaml:",flow"`
	PluginName string   `yaml:"plugin-name"`
	PreInstall string   `yaml:"pre-install"`
}

func (ps *PluginSpec) Encode() *PluginSpec {

	yamlFile, err := ioutil.ReadFile("https://raw.githubusercontent.com/kf5i/k3ai-plugins/main/v2/argo/argo.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, ps)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return ps

}
