package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

var _ Endpoint = (*BaseEndpoint)(nil)

// A BaseEndpoint is an api endpoint.
type BaseEndpoint struct {
	URL          string
	InitFunc     func(verbose bool) error
	ResponseFunc func(req *http.Request) interface{}
}

// GetURL returns the URL where the endpoint is available.
func (ep BaseEndpoint) GetURL() string {
	return ep.URL
}

// Init initializes the endpoint.
// If no InitFunc was provided, does nothing.
func (ep BaseEndpoint) Init(verbose bool) error {
	if ep.InitFunc == nil {
		return nil
	}
	return ep.InitFunc(verbose)
}

// Register registers the endpoint at the http.DefaultServeMux.
func (ep BaseEndpoint) Register() {
	// might not work without suffix because of automatic redirects
	if !strings.HasSuffix(ep.URL, "/") {
		ep.URL += "/" // not permanent
	}

	http.Handle(ep.URL, ep)
}

// ServeHTTP responds to the API request and writes the result to a JSON.
func (ep BaseEndpoint) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	writeJSON(w, ep.ResponseFunc(req))
}

// writeJSON writes an interface as a JSON file to the ResponseWriter.
func writeJSON(w http.ResponseWriter, d interface{}) {

	// for errors, write status code
	if err, ok := d.(Error); ok {
		if err.StatusCode != 0 {
			w.WriteHeader(err.StatusCode)
		}
	}

	// serve with correct MIME type
	w.Header().Set("Content-Type", "application/json")

	// encode json
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t") // make human readable (for now)
	if err := enc.Encode(d); err != nil {
		log.Printf("failed to encode json: %+v", d)
	}
}

// Respond responds to an API request.
func (ep BaseEndpoint) Respond(req *http.Request) interface{} {
	return ep.ResponseFunc(req)
}
