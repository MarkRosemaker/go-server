package form

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"cloud.google.com/go/civil"
)

// some error formats
const (
	notProvided = "%s value not provided"
	notParsed   = "%s value '%s' could not be parsed to %s"
)

func GetStringE(req *http.Request, key string) (string, error) {
	s := req.FormValue(key)
	if s == "" {
		return "", fmt.Errorf(notProvided, key)
	}
	return s, nil
}

func GetString(req *http.Request, key string) string {
	s, _ := GetStringE(req, key)
	return s
}

func GetIntE(req *http.Request, key string) (int, error) {
	s, err := GetStringE(req, key)
	if err != nil {
		return 0, err
	}

	var i int
	i, err = strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf(notParsed, key, s, "int")
	}

	return i, nil
}

func GetInt(req *http.Request, key string) int {
	i, _ := strconv.Atoi(GetString(req, key))
	return i
}

func GetUint64E(req *http.Request, key string) (uint64, error) {
	s, err := GetStringE(req, key)
	if err != nil {
		return 0, err
	}

	var i uint64
	i, err = strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf(notParsed, key, s, "uint64")
	}

	return i, nil
}

func GetUint64(req *http.Request, key string) uint64 {
	i, _ := GetUint64E(req, key)
	return i
}

func GetDurationE(req *http.Request, key string) (time.Duration, error) {
	return time.ParseDuration(req.FormValue(key))
}

func GetDuration(req *http.Request, key string) time.Duration {
	t, _ := GetDurationE(req, key)
	return t
}

func GetBoolE(req *http.Request, key string) (bool, error) {
	s := GetString(req, key)
	if s == "" {
		// string not provided, so default to false
		// there was no error in the sense of "invalid input"
		return false, nil
	}

	// checkbox checked
	if s == "on" {
		return true, nil
	}

	b, err := strconv.ParseBool(s)
	if err != nil {
		// now there was invalid input
		return false, fmt.Errorf(notParsed, key, s, "bool")
	}
	return b, nil
}

func GetBool(req *http.Request, key string) bool {
	b, _ := GetBoolE(req, key)
	return b
}

func GetDateE(req *http.Request, key string) (civil.Date, error) {
	var d civil.Date

	s, err := GetStringE(req, key)
	if err != nil {
		return d, err
	}

	d, err = civil.ParseDate(s)
	if err != nil {
		return d, fmt.Errorf(notParsed, key, s, "date")
	}
	return d, nil
}

func GetDate(req *http.Request, key string) civil.Date {
	d, _ := GetDateE(req, key)
	return d
}
