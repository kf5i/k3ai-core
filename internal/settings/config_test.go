package settings

import (
	"github.com/kf5i/k3ai-core/internal/plugins"
	"github.com/kf5i/k3ai-core/internal/shared"
	"io/ioutil"
	"os"
	"testing"
)

func TestDefaultDirDoesNotExist(t *testing.T) {
	dir, _ := ioutil.TempDir("dir_does_not_exist", "config")

	settingsToRead, err := loadSettingFormFile(dir)
	if err != nil {
		t.Fatalf("can't read setting file, error: %s", err)
	}

	shared.AssertEqual(t, settingsToRead.GroupsURI, plugins.DefaultPluginsGroupURI, "TestDefaultDirDoesNotExist GroupsURI")
	shared.AssertEqual(t, settingsToRead.PluginsURI, plugins.DefaultPluginURI, "TestDefaultDirDoesNotExist PluginsURI")
	shared.AssertEqual(t, settingsToRead.K8sCli, plugins.DefaultK8sCli, "TestDefaultDirDoesNotExist K8sCli")
}

func TestCustomSettings(t *testing.T) {
	dir, err := ioutil.TempDir("", "config")
	if err != nil {
		t.Fatalf("can't save setting file, error: %s", err)
	}
	defer os.RemoveAll(dir)
	var settingsToStore Settings
	settingsToStore.GroupsURI = "path-groups-uri"
	settingsToStore.PluginsURI = "path-plugins-uri"
	settingsToStore.K8sCli = "kubectl"

	err = SaveSettingFile(dir, settingsToStore)
	if err != nil {
		t.Fatalf("can't save setting file, error: %s", err)
	}

	settingsToRead, err := loadSettingFormFile(dir)
	if err != nil {
		t.Fatalf("can't read setting file, error: %s", err)
	}

	shared.AssertEqual(t, settingsToRead.GroupsURI, settingsToStore.GroupsURI, "TestCustomSettings GroupsURI")
	shared.AssertEqual(t, settingsToRead.PluginsURI, settingsToStore.PluginsURI, "TestCustomSettings PluginsURI")
	shared.AssertEqual(t, settingsToRead.K8sCli, settingsToStore.K8sCli, "TestCustomSettings K8sCli")
}

func TestCustomWithEmptySettings(t *testing.T) {
	dir, err := ioutil.TempDir("", "config")
	if err != nil {
		t.Fatalf("can't save setting file, error: %s", err)
	}
	defer os.RemoveAll(dir)
	var settingsToStore Settings
	settingsToStore.GroupsURI = ""
	settingsToStore.PluginsURI = ""
	settingsToStore.K8sCli = ""

	err = SaveSettingFile(dir, settingsToStore)
	if err != nil {
		t.Fatalf("can't save setting file, error: %s", err)
	}

	settingsToRead, err := loadSettingFormFile(dir)
	if err != nil {
		t.Fatalf("can't read setting file, error: %s", err)
	}

	shared.AssertEqual(t, settingsToRead.GroupsURI, plugins.DefaultPluginsGroupURI, "TestCustomWithEmptySettings GroupsURI")
	shared.AssertEqual(t, settingsToRead.PluginsURI, plugins.DefaultPluginURI, "TestCustomWithEmptySettings PluginsURI ")
	shared.AssertEqual(t, settingsToRead.K8sCli, plugins.DefaultK8sCli, "TestCustomWithEmptySettings K8sCli")
}
