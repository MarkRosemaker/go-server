// Package filepath provides conversion of a filepath to a relative URL.
//
// This is necessary on Windows as it uses a different path separator as a URL.
package filepath

import (
	"path/filepath"
	"strings"
)

// ToURL turns a file path into a relative url path.
func ToURL(path string) string {
	// fix seperator on Windows
	return strings.ReplaceAll(filepath.Clean(path), "\\", "/")
}
