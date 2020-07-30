package api

import "net/http"

// Endpoint represents an API endpoint.
type Endpoint interface {
	// GetURL returns the URL where the endpoint is available.
	GetURL() string
	// Init initializes the endpoint.
	Init(verbose bool) error
	// Register registers the endpoint at an http.ServeMux.
	Register()
	// Respond responds to an API request.
	Respond(req *http.Request) interface{}
}
