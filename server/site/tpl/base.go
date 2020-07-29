package tpl

import (
	"html/template"

	"github.com/MarkRosemaker/go-server/server/api"
)

// base template on which other templates are based on
var baseTemplate *template.Template = template.New("")

// the function map
var fm = template.FuncMap{
	"isAPIError": api.IsAPIError,
}

func init() {
	// connect base template to a map of functions
	// note: if we change fm afterwards, it will have no effect on the base template
	baseTemplate.Funcs(fm)
}
