package config

import (
	"html/template"

	"github.com/faizalom/go-web/lib"
)

var ThemePath = "resources/templates/"
var MyTheme = lib.ThemeStruct{}

func init() {

	t := template.Must(template.ParseGlob(ThemePath + "*.html"))
	t.ParseGlob("resources/views/*.html")

	MyTheme = lib.ThemeStruct{
		Template: t,
	}

}
