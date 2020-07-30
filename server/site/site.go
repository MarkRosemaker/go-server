package site

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	_filepath "github.com/MarkRosemaker/go-server/server/filepath"

	"github.com/MarkRosemaker/go-server/server/site/content"
	"github.com/MarkRosemaker/go-server/server/site/content/notfound"
	"github.com/MarkRosemaker/go-server/server/site/tpl"
)

// ErrNoContentSource reports that no content source was provided in the server options.
var ErrNoContentSource error = errors.New("no content source provided")

// InitContent initializes the content from a folder of templates and files.
func InitContent(src string, df tpl.DataFunc) error {
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

			switch info.Name() {
			case "index.html":
				// remove name from url
				url = strings.TrimSuffix(url, info.Name())

				// create page from template
				h, err = content.NewPage(path, url, df)
				if err != nil {
					return err
				}

			case "404.html":
				// remove name from url
				url = strings.TrimSuffix(url, "404.html")

				// create page from template
				h, err = content.NewNotFoundPage(path, url, df)
				if err != nil {
					return err
				}

				// serve as custom 404
				notfound.Handle(url, h)
				return nil
			default:
				// just serve as standard file
				h, err = content.NewFile(path)
				if err != nil {
					return err
				}
			}

			http.Handle(url, h)

			return nil
		})
}
