package data

// Data is an interface representing server data.
type Data interface {
	// Get returns stored data.
	Get(key string) interface{}
	// Load loads the data the server needs.
	Load() error
}

var _ Data = (*None)(nil)

// None is an empty dataset.
type None int

// Get should not be called if we don't have data.
func (n None) Get(key string) interface{} {
	return nil
}

// Load doesn't do anything because we don't have any data.
func (n None) Load() error {
	return nil
}
