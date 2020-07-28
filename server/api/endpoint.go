package api

import "net/http"

// An Endpoint is an api endpoint.
type Endpoint struct {
	// The URL where the endpoint is available.
	URL string
	// HandlerFunc handles an API request.
	http.HandlerFunc
}

// Handle registers the endpoint in the DefaultServeMux.
func (ep Endpoint) Handle() {
	http.HandleFunc(ep.URL, ep.HandlerFunc)
}

// Endpoint implements http.Handler for symmetry's sake

var _ http.Handler = (*Endpoint)(nil)

func (ep Endpoint) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ep.HandlerFunc(w, req)
}
