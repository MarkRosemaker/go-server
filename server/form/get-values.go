package form

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// todo: re-inventing the wheel here?

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

func GetDurationE(req *http.Request, key string) (time.Duration, error) {
	return time.ParseDuration(req.FormValue(key))
}

func GetDuration(req *http.Request, key string) time.Duration {
	t, _ := GetDurationE(req, key)
	return t
}
