package plugins

import (
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

var cacheDirName = ".k3ai/cache"

var errDirTraversalNotAllowed = errors.New("attempted dir traversal when not allowed")

func createCacheDir() error {
	dir, err := cacheDir()
	if err != nil {
		return err
	}
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
	}
	return err
}
func cacheDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, cacheDirName), nil
}

func writeCache(data []byte, uri string) error {
	fullPath, err := uriToFullPath(uri)
	if err != nil {
		return err
	}
	dir := filepath.Dir(fullPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
	}
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fullPath, data, 0644)
}

func readCache(uri string) ([]byte, error) {
	fullPath, err := uriToFullPath(uri)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadFile(fullPath)
}

func cached(uri string) bool {
	fullPath, err := uriToFullPath(uri)
	if err != nil {
		return false
	}
	file, err := os.Stat(fullPath)
	if err != nil {
		return false
	}
	return !file.IsDir()
}

func uriToFullPath(uri string) (string, error) {
	dir, err := cacheDir()
	if err != nil {
		return "", err
	}
	relPath, err := uriToFilePath(uri)
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(dir, relPath)
	// Guard against directory traversal
	if !strings.HasPrefix(fullPath, dir) {
		return "", errDirTraversalNotAllowed
	}
	return filepath.Join(dir, relPath), nil
}

func uriToFilePath(uri string) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	return filepath.Join(u.Host, u.Path), nil
}
