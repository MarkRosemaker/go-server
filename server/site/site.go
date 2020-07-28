package site

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	_filepath "github.com/MarkRosemaker/go-server/server/filepath"

	"github.com/MarkRosemaker/go-server/server/data"
	"github.com/MarkRosemaker/go-server/server/site/content"
)

// ErrNoContentSource reports that no content source was provided in the server options.
var ErrNoContentSource error = errors.New("no content source provided")

// InitContent initializes the content from a folder of templates and files.
func InitContent(src string, d data.Data) error {
	if src == "" {
		return ErrNoContentSource
	}

	// check first if the template folder actually exists (otherwise panic below)
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return fmt.Errorf("can't find template source at %q", src)
	}

	// walk files in tmplSource to find template files
	return filepath.Walk(src,
		func(path string, info os.FileInfo, err error) error {
			// ignore folders
			if info.IsDir() {
				return nil
			}

			url := _filepath.ToURL(path)

			// trim prefix to treat the template folder as its own file system
			url = strings.TrimPrefix(url, src)

			// initialize the page or file
			var h http.Handler

			base := filepath.Base(path)
			name := strings.TrimSuffix(base, filepath.Ext(base))

			switch name {
			case "index":
				// remove base from url
				url = strings.TrimSuffix(url, base)

				// create page from template
				h, err = content.NewPage(path)
				if err != nil {
					return err
				}
			case "404":
				//  serve custom 404 page
				log.Printf("custom 404 pages not implemented yet")
			default:
				// just serve as standard file
				h, err = content.NewFile(path)
				if err != nil {
					return err
				}
			}

			// serve the page when called
			http.Handle(url, h)

			return nil
		})
}
