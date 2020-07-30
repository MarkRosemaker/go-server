package context

import (
	"context"
	"net/http"
	"time"

	"github.com/MarkRosemaker/go-server/server/form"
)

// WithUserTimeout checks the request for a 'timeout' form value.
// If there is a valid one, it returns a background context with this timeout.
// Otherwise, it returns a context without a timeout.
//
// This funciton is meant to be called at the very start of a HandlerFunc.
func WithUserTimeout(req *http.Request) (context.Context, context.CancelFunc) {
	timeout, err := form.GetDurationE(req, "timeout")
	if err == nil {
		return context.WithTimeout(context.Background(), timeout)
	}
	return context.WithCancel(context.Background())
}

// WithUserTimeoutMaxed works like WithUserTimeout, but we add an upper limit to the timout.
//
// This funciton is meant to be called at the very start of a HandlerFunc.
func WithUserTimeoutMaxed(req *http.Request, max time.Duration) (context.Context, context.CancelFunc) {
	timeout, err := form.GetDurationE(req, "timeout")
	if err == nil {
		if timeout > max {
			timeout = max
		}
		return context.WithTimeout(context.Background(), timeout)
	}
	return context.WithCancel(context.Background())
}
