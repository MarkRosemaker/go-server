package notfound

import (
	"net/http"
)

// todo idea: Generalize. Why not have custom 500 pages etc.?

// mux is a server determining which 404 page should be displayed
var mux *http.ServeMux

// Handle adds a custom 404 page via a http.Handler.
//
// Note: The Handler *must not call this package's ServeHTTP*, otherwise we have an infinite loop.
func Handle(pattern string, handler http.Handler) {
	if mux == nil {
		mux = http.NewServeMux()
		serveFunc = mux.ServeHTTP
	}

	mux.Handle(pattern, handler)
}
