package api

import (
	"fmt"
	"time"
)

// Error is an error that can be returned to the user as a json.
type Error struct {
	// when the error occurred
	TimeStamp string `json:"timestamp,omitempty"`

	// HTTP status code
	Status int `json:"status,omitempty"`

	// the http error, such as "Internal Server Error" or "Bad Request"
	HTTPError string `json:"error,omitempty"`

	// the error message for the user
	Message string `json:"message,omitempty"`

	// optional: more details
	Detail string `json:"detail,omitempty"`

	// the url
	Path string `json:"path,omitempty"`

	// e.g.
	// "status":500,
	// "error":"Internal Server Error",
	// "message":"No message available",
	// "path":"/api/book/1"
}

func NewError(status int) Error {
	// HTTPError string `json:"error,omitempty"`
	

	return Error{
		TimeStamp: time.Now().Format(time.RFC3339),
		Status: status,
	}
}

// implement error interface

func (e Error) Error() string {
	return fmt.Sprintf("%+v", e)
}
