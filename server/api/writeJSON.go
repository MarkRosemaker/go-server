package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// WriteJSON writes an interface as a JSON file to the ResponseWriter.
func WriteJSON(w http.ResponseWriter, d interface{}) {
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
