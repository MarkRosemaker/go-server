package api

type Endpoints []Endpoint

// Handle registers all endpoints in the DefaultServeMux.
func (epts Endpoints) Handle() {
	for _, ep := range epts {
		ep.Handle()
	}
}
