package plugins

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// FetchFromSourceURI downloads the content from http or file
func FetchFromSourceURI(uri string) ([]byte, error) {
	if isHTTP(uri) {
		return _fetchRemoteContent(uri)
	}
	return fetchFromFile(uri)

}

func isHTTP(uri string) bool {
	return strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://")
}

func fetchFromFile(uri string) ([]byte, error) {
	fileContent, err := ioutil.ReadFile(uri)
	if err != nil {
		return nil, err
	}
	return fileContent, nil

}

func _fetchRemoteContent(uri string) ([]byte, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	// TODO: Check http status code for better error messages
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func setDefaultIfEmpty(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

func includeSlash(path string) string {
	if strings.HasSuffix(path, "/") {
		return path
	}

	return path + "/"
}

// NormalizePath applies the "/" in the right position
func NormalizePath(args ...string) string {
	result := ""
	for _, subPath := range args {
		result += includeSlash(subPath)
	}
	return strings.TrimRight(result, "/")
}
