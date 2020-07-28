package content

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var _ http.Handler = (*file)(nil)

// A file struct contains variables for quick serving of a file.
type file struct {
	path    string
	modtime time.Time
}

// NewFile creates a new file that is served directly.
func NewFile(path string) (http.Handler, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("couldn't open %q", path)
	}
	defer f.Close()

	d, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("couldn't get stats of %q", path)
	}

	return &file{
		path,
		d.ModTime()}, nil
}

// implement http.Handler

func (fi file) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	f, err := os.Open(fi.path)
	if err != nil {
		// log error for ourselves ...
		log.Printf("couldn't open file %q: %s", fi.path, err)

		// ... and don't leak information in error messages
		switch {
		case os.IsNotExist(err):
			http.Error(w, "404 page not found", http.StatusNotFound)
		case os.IsPermission(err):
			http.Error(w, "403 Forbidden", http.StatusForbidden)
		default:
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	defer f.Close()

	// "The main benefit of ServeContent over io.Copy
	// is that it handles Range requests properly, sets the MIME type, and
	// handles If-Match, If-Unmodified-Since, If-None-Match, If-Modified-Since,
	// and If-Range requests."
	http.ServeContent(w, req, fi.path, fi.modtime, f)
}
