package filepath

import (
	"path/filepath"
	"strings"
)

// ToURL turns a file path into a string.
// An empty string is returned empty so as to not tamper with uninitialized strings.
func ToURL(path string) string {
	if path == "" {
		return ""
	}

	path = filepath.Clean(path)

	// fix seperator on some systems so it has same slashes as url
	if filepath.Separator != '/' {
		return strings.ReplaceAll(path, string(filepath.Separator), "/")
	}

	return path
}
