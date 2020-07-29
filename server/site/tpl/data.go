package tpl

import (
	"net/http"
)

// DataFunc is a function which return is given to a template.
// Since the data for a template depends on the request, that is the input to the function.
type DataFunc func(*http.Request) interface{}

// StdDataFunc is a simple DataFunc that just returns the Request in a struct to the template, i.e. in a template, we can use '{{ .Request }}'.
func StdDataFunc(r *http.Request) interface{} {
	return struct {
		Request *http.Request
	}{r}
}
