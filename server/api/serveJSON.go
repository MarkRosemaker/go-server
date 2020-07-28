package api

import "net/http"

// WriteJSON writes an interface as a JSON file to the ResponseWriter.
func WriteJSON(w http.ResponseWriter, d interface{}) {

	// if d type error:
	// v := make(map[string]string)
	// v["error"]=d.Error()
	// d=v
	// }

	// Write header
	// json ...
}
