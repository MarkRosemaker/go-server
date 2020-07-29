package api

import (
	"encoding/json"
	"fmt"
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

	// todo: would it be helpful?
	// URL string `json:"path,omitempty"`
}

// NewErrorNow creates a new api error message that can be sent to the user.
// It also has a timestamp of the current time.
func NewErrorNow(code int,
	msg, detail string) Error {
	e := NewError(code,
		// url,
		msg, detail)
	e.TimeStamp = time.Now().Format(time.RFC3339)
	return e
}

// NewError creates a new api error message that can be sent to the user.
func NewError(code int,
	// url,
	msg, detail string) Error {
	return Error{
		StatusCode: code,
		StatusText: http.StatusText(code),
		Message:    msg,
		Detail:     detail,
		// URL:        url,
	}
}

// implement error interface

func (e Error) Error() string {
	return fmt.Sprintf("%+v", e)
}

// IsAPIError checks if the input is an api.Error.
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
		// TODO test
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
