package tpl

import (
	"encoding/json"
	"errors"
	"html/template"

	"github.com/spf13/cast"
)

// jsonify encodes a given object to JSON.  To pretty print the JSON, pass a map
// or dictionary of options as the first argument.  Supported options are
// "prefix" and "indent".  Each JSON element in the output will begin on a new
// line beginning with prefix followed by one or more copies of indent according
// to the indentation nesting.
// CREDIT belongs to The Hugo Authors.
func jsonify(args ...interface{}) (template.HTML, error) {
	var (
		b   []byte
		err error
	)

	switch len(args) {
	case 0:
		return "", nil
	case 1:
		b, err = json.Marshal(args[0])
	case 2:
		var opts map[string]string

		opts, err = cast.ToStringMapStringE(args[0])
		if err != nil {
			break
		}

		b, err = json.MarshalIndent(args[1], opts["prefix"], opts["indent"])
	default:
		err = errors.New("too many arguments to jsonify")
	}

	if err != nil {
		return "", err
	}

	return template.HTML(b), nil
}
