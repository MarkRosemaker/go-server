package api

import (
	"fmt"
	"net/http"
	"time"
)

// Success is a success message that can be returned to the user as a json.
type Success struct {
	TimeStamp string `json:"timestamp,omitempty"`

	// HTTP status code
	StatusCode int `json:"code,omitempty"`

	// HTTP status text e.g. "OK", "Created", or "Accepted"
	StatusText string `json:"status,omitempty"`

	// message for the user
	Message string `json:"message,omitempty"`
}

// NewSuccessNow creates a new api success message that can be sent to the user.
// It also has a timestamp of the current time.
func NewSuccessNow(code int, format string, a ...interface{}) Success {
	s := NewSuccess(code, format, a...)
	s.TimeStamp = time.Now().Format(time.RFC3339)
	return s
}

// NewSuccess creates a new api success message that can be sent to the user.
func NewSuccess(code int, format string, a ...interface{}) Success {
	var txt string
	if code != 0 {
		txt = http.StatusText(code)
	}
	return Success{
		StatusCode: code,
		StatusText: txt,
		Message:    fmt.Sprintf(format, a...),
	}
}
