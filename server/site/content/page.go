package content

import (
	"html/template"
	"mime"
	"net/http"
	"path/filepath"

	"github.com/MarkRosemaker/go-server/server/site/tpl"
)

var _ http.Handler = (*page)(nil)

// A page struct contains information for quick serving of a page (or file) from a template.
type page struct {
	tpl      *template.Template
	tplName  string
	mimeType string
}

// NewPage creates a new page (or file) created from a template.
func NewPage(path string) (http.Handler, error) {
	// parse the template
	t, err := tpl.GetTemplate(path)
	if err != nil {
		return nil, err
	}

	return &page{t,
		// store the template name
		filepath.Base(path),
		// get the mime type from the file extension
		mime.TypeByExtension(filepath.Ext(path))}, nil
}

// implement http.Handler

func (p page) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// serve with correct MIME type
	w.Header().Set("Content-Type", p.mimeType)

	// execute the template
	// err := p.tpl.ExecuteTemplate(w, p.tplName, tpldata.New(req))
	// if err != nil {
	// 	log.Printf("couldn't execute template at %q: %s", req.URL.Path, err)
	// 	http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	// }
}
