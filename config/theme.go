package config

import (
	"html/template"

	"github.com/faizalom/go-web/lib"
)

var ThemePath = "resources/templates/"
var MyTheme = lib.ThemeStruct{
	Template: template.Must(template.ParseGlob(ThemePath + "*.html")),
}
