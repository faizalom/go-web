package lib

import (
	"net/http"
	"text/template"

	"github.com/faizalom/go-web/config"
)

var Template = template.Must(template.ParseGlob(config.ThemePath))

// function to run html template
func ExeTemplate(w http.ResponseWriter, templateFile string, data any) {
	err := Template.ExecuteTemplate(w, templateFile, data)
	if err != nil {
		panic(err)
	}
}
