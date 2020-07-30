package content

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MarkRosemaker/go-server/server/site/tpl"
)

// todo idea: Generalize to a problemPage. Why not have custom 500 pages etc.?

type notFoundPage struct {
	*page
}

// NewNotFoundPage returns a new page that can be shown if another page can't be found.
func NewNotFoundPage(path, url string, data tpl.DataFunc) (http.Handler, error) {
	p, err := newPage(path, url, data)
	if err != nil {
		return nil, err
	}

	return &notFoundPage{p}, nil
}

// implement http.Handler

func (nf notFoundPage) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	if err := nf.executeTemplate(w, req); err != nil {

		// do what's essentially http.Error with the standard line
		// but don't write header again
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		fmt.Fprintln(w, "404 page not found")

		// log error
		log.Printf("404 template: %s", err)
	}
}
