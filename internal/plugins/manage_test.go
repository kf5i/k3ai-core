package plugins

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func mockServer(t *testing.T) *httptest.Server {
	file := getTestSpecFile(t, "defaults/plugin.yaml")
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasSuffix(r.URL.Path, "/argo") {
				fmt.Fprintln(w, mockArgoFolder(r.Host))
			} else if strings.HasSuffix(r.URL.Path, ".yaml") {
				fmt.Fprintln(w, string(file))
			} else {
				fmt.Fprintln(w, pluginList)
			}
		}))
	return ts
}

func TestPluginList(t *testing.T) {
	var server = mockServer(t)
	defer server.Close()
	p, err := PluginList(server.URL)
	if err != nil {
		t.Fatalf("expected nil but got %v", err)
	}
	if 2 != len(p) {
		t.Fatalf("expected %d but got %v", 2, len(p))
	}
}

func TestPluginYamls(t *testing.T) {
	var server = mockServer(t)
	defer server.Close()
	p, err := PluginYamls(server.URL+"/", "argo")
	if err != nil {
		t.Fatalf("expected nil but got %v", err)
	}
	if 1 != len(p) {
		t.Fatalf("expected %d but got %v", 1, len(p))
	}
}

const pluginList = `[
			{
				"name": "README.md",
				"download_url": "https://raw.githubusercontent.com/kf5i/k3ai-plugins/main/v2/README.md",
				"type": "file"
			},
			{
				"name": "argo",
				"download_url": null,
				"type": "dir"
			},
			{
				"name": "tensorflow",
				"download_url": null,
				"type": "dir"
			}
		]`

func mockArgoFolder(serverURL string) string {
	return fmt.Sprintf(`[
  {
		"name": "plugin.yaml",
		"type": "file",
    "download_url": "http://%s/argo/plugin.yaml"
	},
	{
				"name": "README.md",
				"download_url": "http://%s/v2/argo/README.md",
				"type": "file"
	}
]`, serverURL, serverURL)
}
