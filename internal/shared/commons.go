package shared

import (
	"os"
	"strings"
)

// IncludeSlash append the / where needed
func IncludeSlash(path string, typeSeparator string) string {
	if strings.HasSuffix(path, typeSeparator) {
		return path
	}
	return path + typeSeparator
}

//IncludeOsSeparator include os path separator
func IncludeOsSeparator(path string) string {
	return IncludeSlash(path, string(os.PathSeparator))
}

// NormalizePath applies the "/" in the right position
func NormalizePath(file string, args ...string) string {
	result := ""
	for _, subPath := range args {
		result += IncludeOsSeparator(subPath)
	}
	return result + file
}

// NormalizeURL applies the "/" in the right position
func NormalizeURL(args ...string) string {
	result := ""
	for _, subPath := range args {
		result += IncludeSlash(subPath, "/")
	}
	return result
}

// GetDefaultIfEmpty get a value if empty
func GetDefaultIfEmpty(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
