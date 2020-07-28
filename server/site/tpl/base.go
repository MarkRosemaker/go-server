package tpl

import "html/template"

// base template on which other templates are based on
var baseTemplate *template.Template = template.New("")

func init() {
	// connect base template to a map of functions
	// note: if we change fm afterwards, it will have no effect on the base template
	baseTemplate.Funcs(fm)
}
