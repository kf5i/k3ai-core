package shared

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"strings"

	"gopkg.in/yaml.v3"
)

//Config is the constructor to read the infra config file
type Config struct {
	Kind                string          `yaml:"kind"`
	TargetCustomization []TargetCustoms `yaml:"targetCustomizations,flow"`
}

//PluginGroup represent the structure for the inline plugins
type PluginGroup struct {
	Repo string `yaml:"repo,omitempty"`
	Name string `yaml:"name,omitempty"`
}

//TargetCustoms represent the single customization group
type TargetCustoms struct {
	Name              string `yaml:"name"`
	Enabled           bool   `yaml:"enabled"`
	Type              string `yaml:"type"`
	Config            string `yaml:"config"`
	ClusterName       string `yaml:"clusterName"`
	ClusterDeployment string `yaml:"clusterDeployment"`
	ClusterStart      string `yaml:"clusterStart,omitempty"`
	Spec              struct {
		Wsl             string `yaml:"wsl,omitempty"`
		Mac             string `yaml:"mac,omitempty"`
		Linux           string `yaml:"linux,omitempty"`
		Windows         string `yaml:"windows,omitempty"`
		Arm         	string `yaml:"arm,omitempty"`
		cloudType       string `yaml:"cloudType,omitempty"`
		cloudNodes      string `yaml:"cloudNodes,omitempty"`
		cloudSecretPath string `yaml:"cloudSecretPath,omitempty"`
	}
	Plugins []PluginGroup `yaml:"plugins,flow"`
}

//Init read the config file and return the json string
func Init(data string) (cfg Config, err error) {
	usr, _ := user.Current()

	if data == "/.k3ai/config.yaml" {
		data = usr.HomeDir + data
	}

	yamlFile, err := ioutil.ReadFile(data)
	if err != nil {
		//fmt.Printf("Error reading YAML file: %s\n", err)
		// result := strings.SplitN(data, "/", 3)
		// let's split dir from file
		// var homeDirStr= usr.HomeDir + "/"+ string(result[1])
		var homeDirStr = usr.HomeDir + "/" + ".k3ai/"
		if _, err := os.Stat(homeDirStr); os.IsNotExist(err) {
			os.Mkdir(homeDirStr, 0777)
			fileURL := "https://raw.githubusercontent.com/kf5i/k3ai-core/main/configs/config.yaml"
			err := downloadFile(data, fileURL)
			if err != nil {
				panic(err)
			}
			//fmt.Println("Downloaded: " + fileURL)
		} else {
			fileURL := "https://raw.githubusercontent.com/kf5i/k3ai-core/main/configs/config.yaml"
			err := downloadFile(data, fileURL)
			if err != nil {
				panic(err)
			}
		}

	}
	yamlFile = []byte(yamlFile)
	yamlDecode := yaml.NewDecoder(bytes.NewReader(yamlFile))
	for {
		var doc Config
		if yamlDecode.Decode(&doc) != nil {
			break
		}
		return doc, nil
	}
	return
}

func createConfig(data string) error {
	usr, _ := user.Current()
	result := strings.SplitN(data, "/", 3)
	// let's split dir from file
	var homeDirStr = usr.HomeDir + "/" + string(result[1])
	if _, err := os.Stat(homeDirStr); os.IsNotExist(err) {
		os.Mkdir(homeDirStr, 600)
		fileURL := "https://raw.githubusercontent.com/kf5i/k3ai-core/main/configs/config.yaml"
		err := downloadFile(data, fileURL)
		if err != nil {
			panic(err)
		}
		fmt.Println("Downloaded: " + fileURL)
	}

	return nil
}

func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
