package tpl

import (
	"html/template"
	"time"

	"cloud.google.com/go/civil"

	"github.com/MarkRosemaker/go-server/server/api"
	"github.com/spf13/cast"
)

// base template on which other templates are based on
var baseTemplate *template.Template = template.New("")

// the function map
var fm = template.FuncMap{
	"isAPIError": api.IsAPIError,
	"dateFormat": dateFormat,
}

func init() {
	// connect base template to a map of functions
	// note: if we change fm afterwards, it will have no effect on the base template
	baseTemplate.Funcs(fm)
}

// from Hugo (github.com/gohugoio/hugo/tpl/time/time.go: Format):
// dateFormat converts the textual representation of the datetime string into
// the other form or returns it of the time.Time value. These are formatted
// with the layout string
func dateFormat(layout string, v interface{}) (string, error) {
	var t time.Time
	switch d := v.(type) {
	case civil.Date:
		t = d.In(time.Local)
	case time.Time:
		t = d
	default:
		var err error
		t, err = cast.ToTimeE(v)
		if err != nil {
			return "", err
		}
	}

	return t.Format(layout), nil
}
