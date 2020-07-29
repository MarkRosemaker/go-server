package server

import (
	"log"
	"net/http"

	"github.com/MarkRosemaker/go-server/server/site"
)

// Run runs the server with the given options.
func Run(o Options) {
	if o.Verbose {
		log.Printf("initializing server")
	}

	// "clean" options
	o.resolve()

	// // load server data
	// if err := o.Data.Load(); err != nil {
	// 	log.Fatalf("error loading server data: %s", err)
	// }

	// register the API endpoints
	o.Endpoints.Handle()

	// initialize the content of the site
	if err := site.InitContent(o.ContentSource, o.TemplateDataFunc); err != nil {
		if err != site.ErrNoContentSource {
			log.Fatalf("error initializing content: %s", err)
		}
		// log ErrNoContentSource, maybe we just want to have API endpoints
		log.Println(err)
	}

	// start the server
	if o.Verbose {
		log.Printf("starting server")
	}
	panic(http.ListenAndServe(o.Address, nil))
}
