package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MarkRosemaker/go-server/server/site"
)

// Run runs the server with the given options.
func Run(o Options) {
	if o.Verbose {
		log.Println("initializing server")
	}

	// "clean" options
	o.resolve()

	// initilizing endpoints, e.g. loading necessary data
	if err := o.Endpoints.Init(o.Verbose, o.EndpointsInitConcurrently); err != nil {
		log.Fatalf("error initializing endpoints: %s", err)
	}

	// register the API endpoints
	o.Endpoints.Register()

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
		log.Println("starting server")
		fmt.Printf("Visit the server at http://localhost%s.\n", o.Address)
	}
	panic(http.ListenAndServe(o.Address, nil))
}
