package api

import (
	"golang.org/x/sync/errgroup"
)

// Endpoints are a slice of API endpoints.
type Endpoints []Endpoint

// Init initializes all endpoints, either concurrently or one after the other.
// It returns an error if one endpoint failed to initialize.
func (epts Endpoints) Init(verbose, conc bool) error {
	if conc {
		errg := new(errgroup.Group)
		for _, ep := range epts {
			errg.Go(func() error {
				return ep.Init(verbose)
			})
		}
		return errg.Wait()
	}

	for _, ep := range epts {
		if err := ep.Init(verbose); err != nil {
			return err
		}
	}

	return nil
}

// Register registers all endpoints in their ServeMux.
func (epts Endpoints) Register() {
	for _, ep := range epts {
		ep.Register()
	}
}
