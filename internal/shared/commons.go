package shared

import (
	"os"
	"strings"
)

// IncludeSlash append the / where needed
func IncludeSlash(path string) string {
	if strings.HasSuffix(path, string(os.PathSeparator)) {
		return path
	}
	return path + string(os.PathSeparator)
}

// NormalizePath applies the "/" in the right position
func NormalizePath(file string, args ...string) string {
	result := ""
	for _, subPath := range args {
		result += IncludeSlash(subPath)
	}
	return result + file
}
