package plugins

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func mockPluginsServer(t *testing.T, filePath string, contentType string) *httptest.Server {
	file, err := fetchFromFile(filePath)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasSuffix(r.URL.Path, "/test") {
				_, _ = fmt.Fprintln(w, mockTestFolder(r.Host, contentType))
			} else if strings.HasSuffix(r.URL.Path, ".yaml") {
				_, _ = fmt.Fprintln(w, string(file))
			} else {
				_, _ = fmt.Fprintln(w, remoteDir)
			}
		}))
	return ts
}

func TestPluginsList(t *testing.T) {
	var server = mockPluginsServer(t, joinWithRootData("plugins/argo-workflow-ns/plugin.yaml"), PluginType)
	defer server.Close()
	p, err := ContentList(server.URL)
	if err != nil {
		t.Fatalf("expected nil but got %v", err)
	}
	if 2 != len(p) {
		t.Fatalf("expected %d but got %v", 2, len(p))
	}
}

func TestPluginYamls(t *testing.T) {
	var server = mockPluginsServer(t, joinWithRootData("plugins/argo-workflow-ns/plugin.yaml"), PluginType)
	defer server.Close()
	var pluginList Plugins

	p, err := pluginList.Encode(server.URL, "/test")

	if err != nil {
		t.Fatalf("expected nil but got %v", err)
	}
	if 1 != len(p.Plugins) {
		t.Fatalf("expected %d but got %v", 1, len(p.Plugins))
	}
}

func TestGroupsYamls(t *testing.T) {
	var server = mockPluginsServer(t, joinWithRootData("groups/argo/group.yaml"), GroupType)
	defer server.Close()
	var groups Groups
	r, err := groups.Encode(server.URL, "/test")

	if err != nil {
		t.Fatalf("expected nil but got %v", err)
	}

	if 1 != len(r.Groups) {
		t.Fatalf("expected %d but got %v", 1, len(r.Groups))
	}
}

const remoteDir = `[
			{
				"name": "README.md",
				"download_url": "https://raw.githubusercontent.com/kf5i/k3ai-plugins/main/core/README.md",
				"type": "file"
			},
			{
				"name": "test",
				"download_url": null,
				"type": "dir"
			},
			{
				"name": "tensorflow",
				"download_url": null,
				"type": "dir"
			}
		]`

func mockTestFolder(serverURL string, contentType string) string {
	s := fmt.Sprintf(`[
  {
		"name": "%s.yaml",
		"type": "file",
	    "download_url": "http://%s/test/%s.yaml"
	},
	{
	     "name": "README.md",
         "download_url": "http://%s/core/%s/test/README.md",
         "type": "file"
	}
]`, contentType, serverURL, contentType, serverURL, contentType)
	return s
}
