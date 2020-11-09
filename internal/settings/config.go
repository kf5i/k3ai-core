package settings

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/kf5i/k3ai-core/internal/plugins"
	"github.com/kf5i/k3ai-core/internal/shared"
	"gopkg.in/yaml.v2"
)

// Settings default user setting
type Settings struct {
	// PluginsURI is --plugin-repo
	PluginsURI string `yaml:"plugins-url"`
	// GroupsURI is --group-repo
	GroupsURI string `yaml:"groups-url"`
	// UseKubectl use kubectl instead of k3s
	UseKubectl bool `yaml:"use-kubectl"`
}

const configFileName = "config"

// GetDefaultSettings get default settings
func GetDefaultSettings() *Settings {
	var ds Settings
	ds.GroupsURI = plugins.DefaultPluginsGroupURI
	ds.PluginsURI = plugins.DefaultPluginURI
	ds.UseKubectl = false
	return &ds
}

// SaveSettingFileHome save the setting to the home
func SaveSettingFileHome(settings Settings) error {
	home, err := getHomeDir()
	if err != nil {
		return err
	}
	return SaveSettingFile(home, settings)
}

// SaveSettingFile save the setting to a generic path
func SaveSettingFile(configDir string, settings Settings) error {

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		errDir := os.MkdirAll(configDir, 0755)
		if errDir != nil {
			return errDir
		}
	}
	err := createSettingFile(configDir)
	if err != nil {
		return err
	}

	yamlContent, err := yaml.Marshal(&settings)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(configDir, configFileName), yamlContent, 0644)

	if err != nil {
		return err
	}
	return nil
}

func createSettingFile(configDir string) error {
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

func getHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	configDir := shared.IncludeSlash(usr.HomeDir) + ".k3ai"
	return configDir, nil
}

// loadSettingFormFile load the configuration file
func loadSettingFormFile(configDir string) (*Settings, error) {
	var ds = GetDefaultSettings()
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		return ds, nil
	}
	if _, err := os.Stat(filepath.Join(configDir, configFileName)); os.IsNotExist(err) {
		return ds, nil
	}

	data, err := ioutil.ReadFile(filepath.Join(configDir, configFileName))
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &ds)
	if err != nil {
		return nil, err
	}

	ds.GroupsURI = shared.GetDefaultIfEmpty(ds.GroupsURI, plugins.DefaultPluginsGroupURI)
	ds.PluginsURI = shared.GetDefaultIfEmpty(ds.PluginsURI, plugins.DefaultPluginURI)

	return ds, nil
}

// LoadSettingFormHomeFile global function to get the settings from home directory
func LoadSettingFormHomeFile() (*Settings, error) {
	home, err := getHomeDir()
	if err != nil {
		return nil, err
	}
	ds, err := loadSettingFormFile(home)
	if err != nil {
		log.Printf("Can't load the configuration, errror: %s", err)
		return GetDefaultSettings(), nil
	}

	return ds, nil

}
