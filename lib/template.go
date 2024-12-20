package lib

import (
	"net/http"
	"text/template"
)

var Template *template.Template

func TemplateParseGlob(ThemePath string) {
	Template = template.Must(template.ParseGlob(ThemePath))
}

// function to run html template
func ExeTemplate(w http.ResponseWriter, templateFile string, data any) {
	err := Template.ExecuteTemplate(w, templateFile, data)
	if err != nil {
		panic(err)
	}
}
