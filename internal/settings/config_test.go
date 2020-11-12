package settings

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/kf5i/k3ai-core/internal/plugins"
	"github.com/kf5i/k3ai-core/internal/shared"
)

func TestDefaultDirDoesNotExist(t *testing.T) {
	dir, _ := ioutil.TempDir("dir_does_not_exist", "config")

	settingsToRead, err := loadSettingFormFile(dir)
	if err != nil {
		t.Fatalf("can't read setting file, error: %s", err)
	}

	shared.AssertEqual(t, settingsToRead.GroupRepo, plugins.DefaultPluginsGroupURI, "TestDefaultDirDoesNotExist GroupsURI")
	shared.AssertEqual(t, settingsToRead.PluginRepo, plugins.DefaultPluginURI, "TestDefaultDirDoesNotExist PluginsURI")
	shared.AssertEqual(t, settingsToRead.UseKubectl, false, "TestDefaultDirDoesNotExist K8sCli")
}

func TestCustomSettings(t *testing.T) {
	dir, err := ioutil.TempDir("", "config")
	if err != nil {
		t.Fatalf("can't save setting file, error: %s", err)
	}
	defer os.RemoveAll(dir)
	var settingsToStore Settings
	settingsToStore.GroupRepo = "path-groups-uri"
	settingsToStore.PluginRepo = "path-plugins-uri"
	settingsToStore.UseKubectl = true

	err = SaveSettingFile(dir, settingsToStore)
	if err != nil {
		t.Fatalf("can't save setting file, error: %s", err)
	}

	settingsToRead, err := loadSettingFormFile(dir)
	if err != nil {
		t.Fatalf("can't read setting file, error: %s", err)
	}

	shared.AssertEqual(t, settingsToRead.GroupRepo, settingsToStore.GroupRepo, "TestCustomSettings GroupsURI")
	shared.AssertEqual(t, settingsToRead.PluginRepo, settingsToStore.PluginRepo, "TestCustomSettings PluginsURI")
	shared.AssertEqual(t, settingsToRead.UseKubectl, settingsToStore.UseKubectl, "TestCustomSettings K8sCli")
}

func TestCustomWithEmptySettings(t *testing.T) {
	dir, err := ioutil.TempDir("", "config")
	if err != nil {
		t.Fatalf("can't save setting file, error: %s", err)
	}
	defer os.RemoveAll(dir)
	var settingsToStore Settings
	settingsToStore.GroupRepo = ""
	settingsToStore.PluginRepo = ""
	settingsToStore.UseKubectl = false

	err = SaveSettingFile(dir, settingsToStore)
	if err != nil {
		t.Fatalf("can't save setting file, error: %s", err)
	}

	settingsToRead, err := loadSettingFormFile(dir)
	if err != nil {
		t.Fatalf("can't read setting file, error: %s", err)
	}

	shared.AssertEqual(t, settingsToRead.GroupRepo, plugins.DefaultPluginsGroupURI, "TestCustomWithEmptySettings GroupsURI")
	shared.AssertEqual(t, settingsToRead.PluginRepo, plugins.DefaultPluginURI, "TestCustomWithEmptySettings PluginsURI ")
	shared.AssertEqual(t, settingsToRead.UseKubectl, false, "TestCustomWithEmptySettings K8sCli")
}
