package server

import (
	"github.com/MarkRosemaker/go-server/server/api"
	"github.com/MarkRosemaker/go-server/server/filepath"
	"github.com/MarkRosemaker/go-server/server/site/tpl"
)

// Options is a struct holding options for our server.
type Options struct {
	// The TCP network address, defaults to ":8080".
	Address string
	// The source to the website content.
	ContentSource string
	// The data given to a template as a function of the request.
	TemplateDataFunc tpl.DataFunc
	// API Endpoints we want to serve.
	Endpoints api.Endpoints
	// Set to true if you want to initialize the endpoints concurrently.
	EndpointsInitConcurrently bool
	// Whether we want to log verbosely.
	Verbose bool
}

// resolve will make sure every field of Options is initialized by setting uninitialized fields to default values.
func (o *Options) resolve() {
	o.ContentSource = filepath.ToURL(o.ContentSource)

	if o.Address == "" {
		o.Address = ":8080"
	}

	if o.TemplateDataFunc == nil {
		o.TemplateDataFunc = tpl.StdDataFunc
	}

	if o.Endpoints == nil {
		o.Endpoints = make(api.Endpoints, 0)
	}
}
