package plugins

import "github.com/kf5i/k3ai-core/internal/shared"

type CacheItem struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type Cache struct {
	Name  string  `yaml:"name"`
	Type  string  `yaml:"type"`
	Items []CacheItem `yaml:"items,flow"`
}

func (cache *Cache) Encode(URL string, fileCache string) error {
	cacheUrl := shared.IncludeSlash(URL, "/") + fileCache
	return encode(cacheUrl, cache)
}
