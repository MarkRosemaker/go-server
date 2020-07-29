package notfound

import (
	"net/http"
)

var serveFunc http.HandlerFunc = http.NotFound

// ServeHTTP serves a not found page.
func ServeHTTP(w http.ResponseWriter, req *http.Request) {
	serveFunc(w, req)
	log(req.URL)
}
