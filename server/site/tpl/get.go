// Package tpl defines behavior that relates to templates.
// Since this is not really the point of the exercise, the implementation is very rudimentary.
package tpl

import (
	"fmt"
	"html/template"
)

// IsTemplate checks wheather the file is a temlpate or not.
func IsTemplate(path string) bool {
	// TODO

	return true
}

// GetTemplate clones the base template and adds a parsed template from the path.
func GetTemplate(path string) (*template.Template, error) {
	// clone the base template
	t, err := baseTemplate.Clone()
	if err != nil {
		return nil, fmt.Errorf("error cloning base template: %s", err)
	}

	// parse the file
	if t, err = t.ParseFiles(path); err != nil {
		return nil, fmt.Errorf("error parsing template at %q: %s", path, err)
	}

	return t, nil
}
