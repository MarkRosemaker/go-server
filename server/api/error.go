package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// Error is an error that can be returned to the user as a json.
type Error struct {
	// when the error occurred
	TimeStamp string `json:"timestamp,omitempty"`

	// HTTP status code
	StatusCode int `json:"status,omitempty"`

	// HTTP status text e.g. "Internal Server Error" or "Bad Request"
	StatusText string `json:"error,omitempty"`

	// error message for the user
	Message string `json:"message,omitempty"`

	// optional
	Detail string `json:"detail,omitempty"`
}

// NewErrorNow creates a new api error message that can be sent to the user.
// It also has a timestamp of the current time.
func NewErrorNow(code int,
	msg, detail string) Error {
	e := NewError(code, msg, detail)
	e.TimeStamp = time.Now().Format(time.RFC3339)
	return e
}

// NewError creates a new api error message that can be sent to the user.
func NewError(code int,	msg, detail string) Error {
	return Error{
		StatusCode: code,
		StatusText: http.StatusText(code),
		Message:    msg,
		Detail:     detail,
	}
}

// implement error interface

func (e Error) Error() string {
	if e.Message == "" {
		return fmt.Sprintf("%d %s", e.StatusCode, e.StatusText)
	}
	return fmt.Sprintf("%d %s: %s", e.StatusCode, e.StatusText, e.Message)
}

// IsAPIError checks if the input is an api.Error for use in templates.
//
// It also returns true if it's the JSON represenation of one,
// so we can use this function in templates.
func IsAPIError(d interface{}) bool {
	switch s := d.(type) {
	case Error:
		return true
	case string:
		if s == "" {
			return false
		}
		// todo test
		var e Error
		dec := json.NewDecoder(strings.NewReader(s))
		err := dec.Decode(&e)
		if err != nil {
			return false
		}
		// empty json would also not be an error
		return e != Error{}
	default:
		return false
	}
}

// ErrBadRequest returns an api.Error with status bad request and the error message.
func ErrBadRequest(err error) Error {
	return NewErrorNow(
		http.StatusBadRequest,
		err.Error(),
		"")
}

// ErrWrap wraps an eror into an api.Error struct.
func ErrWrap(err error) Error {
	// is already wrapped
	if w, ok := err.(Error); ok {
		return w
	}

	if err == context.DeadlineExceeded {
		// "The 408 (Request Timeout) status code indicates that the server did not receive a complete request message within the time that it was prepared to wait."
		return NewErrorNow(http.StatusRequestTimeout, "", "")
	}

	// log but don't leak information
	log.Printf("api.ErrWrap called to wrap error: %s", err)
	return NewErrorNow(http.StatusInternalServerError, "unknown error", "")
}
