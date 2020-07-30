package content

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"mime"
	"net/http"
	"path/filepath"

	"github.com/MarkRosemaker/go-server/server/site/content/notfound"
	"github.com/MarkRosemaker/go-server/server/site/tpl"
)

var _ http.Handler = (*page)(nil)

// A page struct contains information for quick serving of a page (or file) from a template.
type page struct {
	tpl      *template.Template
	tplName  string
	url      string
	data     tpl.DataFunc
	mimeType string
}

// NewPage creates a new page (or file) created from a template.
func NewPage(path, url string, data tpl.DataFunc) (http.Handler, error) {
	return newPage(path, url, data)
}

// newPage is an internal function to create a new page (or file) created from a template.
// The return is a pointer to a page struct (leaving it exposed) and an error.
func newPage(path, url string, data tpl.DataFunc) (*page, error) {
	// parse the template
	t, err := tpl.GetTemplate(path)
	if err != nil {
		return nil, err
	}

	return &page{
		// store the template
		t,
		// store the template name
		filepath.Base(path),
		// store the url
		url,
		// store the data function
		data,
		// get the mime type from the file extension
		mime.TypeByExtension(filepath.Ext(path))}, nil
}

// implement http.Handler

func (p page) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// this needs to be checked in case the request url is longer
	// i.e. /foo/bar/ instead of /foo/, /foo/bar/ should yield 404
	if req.URL.Path != p.url {
		notfound.ServeHTTP(w, req)
		return
	}

	if err := p.executeTemplate(w, req); err != nil {

		// log error and give general feedback to user
		log.Println(err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	// serve with correct MIME type
	w.Header().Set("Content-Type", p.mimeType)
}

func (p page) executeTemplate(w http.ResponseWriter, req *http.Request) error {
	// execute the template into a buffer
	// otherwise, we run into problems if template is not executed
	// todo: explore effects on performance; we should write our templates in such a way that they don't break in the first place, after all
	// idea: maybe decide according to which phase of development we're in, the binary could be compiled using certain build flags

	var b bytes.Buffer
	err := p.tpl.ExecuteTemplate(&b, p.tplName, p.data(req))
	if err != nil {
		return fmt.Errorf("couldn't execute template at %q: %s", req.URL.Path, err)
	}
	w.Write(b.Bytes())
	return nil
}
