package plugins

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	// GroupType  Group Type
	GroupType = "group"
	// PluginType Plugin Type
	PluginType = "plugin"
)

// FetchFromSourceURI downloads the content from http or file
func FetchFromSourceURI(uri string) ([]byte, error) {
	if isHTTP(uri) {
		return fetchRemoteContent(uri)
	}
	return fetchFromFile(uri)

}

func isHTTP(uri string) bool {
	return strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://")
}

// fetchFromFile load the yaml from file
func fetchFromFile(uri string) ([]byte, error) {
	fileContent, err := ioutil.ReadFile(uri)
	if err != nil {
		return nil, err
	}
	return fileContent, nil

}

func fetchRemoteContent(uri string) ([]byte, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	// TODO: Check http status code for better error messages
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func getDefaultIfEmpty(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
