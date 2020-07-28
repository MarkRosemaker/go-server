package tpl

import (
	"html/template"
)

// the function map
var fm = template.FuncMap{
	"jsonify": jsonify,
	// "login":    auth.Login,
	// "signup":   auth.Signup,
}
