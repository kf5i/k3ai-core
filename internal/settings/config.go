package settings

import (
	"github.com/ghodss/yaml"
	"github.com/kf5i/k3ai-core/internal/plugins"
	"github.com/kf5i/k3ai-core/internal/shared"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

// DefaultSettings default user setting
type DefaultSettings struct {
	// PluginsURI is --plugin-repo
	PluginsURI string `yaml:"plugins-url"`
	// GroupsURI is --group-repo
	GroupsURI string `yaml:"groups-url"`
	// K8sCli is k3s or kubectl
	K8sCli string `yaml:"k8s-cli"`
}

const configFileName = "config"

func GetDefaultSettings() *DefaultSettings {
	var ds DefaultSettings
	ds.GroupsURI = plugins.DefaultPluginsGroupURI
	ds.PluginsURI = plugins.DefaultPluginURI
	ds.K8sCli = "k3s"
	return &ds
}

// SaveConfigurationFile save the config
func (ds *DefaultSettings) SaveConfigurationFile() error {
	configDir, err := getHomeConfigDir()
	if _, err = os.Stat(configDir); os.IsNotExist(err) {
		errDir := os.MkdirAll(configDir, 0755)
		if errDir != nil {
			return errDir
		}
	}
	err = CreateConfigFile(configDir)
	if err != nil {
		return err
	}

	yamlContent, err := yaml.Marshal(ds)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(configDir, configFileName), yamlContent, 0644)

	if err != nil {
		return nil
	}
	return nil
}

func CreateConfigFile(configDir string) error {
	if _, err := os.Stat(filepath.Join(configDir, configFileName)); os.IsNotExist(err) {
		f, err := os.Create(filepath.Join(configDir, configFileName))
		if err != nil {
			return err
		}
		err = f.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func getHomeConfigDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	configDir := shared.IncludeSlash(usr.HomeDir) + ".k3ai"
	return configDir, nil
}

func LoadConfigFile() (*DefaultSettings, error) {
	var ds *DefaultSettings
	configDir, err := getHomeConfigDir()
	if _, err = os.Stat(configDir); os.IsNotExist(err) {
		return GetDefaultSettings(), nil
	}
	if _, err := os.Stat(filepath.Join(configDir, configFileName)); os.IsNotExist(err) {
		return GetDefaultSettings(), nil

	}

	data, err := ioutil.ReadFile(filepath.Join(configDir, configFileName))
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, ds)
	if err != nil {
		return nil, err
	}

	return ds, nil
}
