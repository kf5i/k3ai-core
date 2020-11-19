package plugins

import (
	"fmt"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"

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

	data, err := readCache(uri)
	// Ignore cache read error
	if err == nil {
		fmt.Println("cache hit")
		return data, nil
	}

	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	// TODO: Check http status code for better error messages
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	if err == nil {
		// Ignore cache write issue
		e := writeCache(data, uri)
		fmt.Println(e)
	}
	return data, err
}

func encode(URL string, value interface{}) error {
	downloadURL := URL
	if isHTTP(URL) {
		gHubContent, err := githubContent(URL)
		if err != nil {
			return errors.Wrap(err, "error fetching plugins content")
		}
		downloadURL = gHubContent.DownloadURL
	}

	remoteContent, err := FetchFromSourceURI(downloadURL)
	if err != nil {
		return errors.Wrap(err, "error fetching plugins group spec")
	}
	err = yaml.Unmarshal(remoteContent, value)
	if err != nil {
		return err
	}
	return nil
}
