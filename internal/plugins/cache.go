package plugins

import "github.com/kf5i/k3ai-core/internal/shared"

// CacheItem the item cache
type CacheItem struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

// Cache for plugins and groups
type Cache struct {
	Name  string      `yaml:"name"`
	Type  string      `yaml:"type"`
	Items []CacheItem `yaml:"items,flow"`
}

// Encode encode the cache to be printed
func (cache *Cache) Encode(URL string, fileCache string) error {
	cacheURL := shared.IncludeSlash(URL, "/") + fileCache
	return encode(cacheURL, cache)
}
