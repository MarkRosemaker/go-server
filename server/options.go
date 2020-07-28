package server

import (
	"github.com/MarkRosemaker/go-server/server/api"
	"github.com/MarkRosemaker/go-server/server/data"
	"github.com/MarkRosemaker/go-server/server/filepath"
)

// Options is a struct holding options for our server.
type Options struct {
	// The TCP network address, defaults to ":8080".
	Address string
	// The source to the website content.
	ContentSource string
	// The server data.
	data.Data
	// API Endpoints we want to serve.
	api.Endpoints
	// Whether we want to log verbosely.
	Verbose bool
}

// resolve will make sure every field of Options is initialized by setting uninitialized fields to default values.
func (o *Options) resolve() {
	o.ContentSource = filepath.ToURL(o.ContentSource)

	if o.Address == "" {
		o.Address = ":8080"
	}

	if o.Data == nil {
		o.Data = new(data.None)
	}

	if o.Endpoints == nil {
		o.Endpoints = make(api.Endpoints, 0)
	}
}
