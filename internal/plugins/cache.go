package plugins

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
)

const cacheDirName = ".k3ai/cache"

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
	// Has not been tested on windows
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
	fmt.Println(dir)
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
	return filepath.Join(dir, relPath), nil
}

func uriToFilePath(uri string) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	return u.Path, nil
}
