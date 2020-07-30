// Package filepath provides conversion of a filepath to a relative URL.
//
// Unix systems already have the same path separator as a URL, so ToURL just needs to clean the path.

// +build aix darwin dragonfly freebsd js,wasm linux nacl netbsd openbsd solaris

package filepath

import (
	"path/filepath"
)

// ToURL turns a file path into a relative url path.
func ToURL(path string) string {
	return filepath.Clean(path)
}
