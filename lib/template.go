package lib

import (
	"html/template"
	"net/http"

	"github.com/faizalom/go-web/models"
	"github.com/gorilla/csrf"
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

// function to run html template
func (a AppStruct) ExeTemp(w http.ResponseWriter, r *http.Request, templateFile string, data any) {
	var res = struct {
		Title       string
		Flash       map[string]any
		CurrentPath string
		AuthUser    models.User
		CsrfField   template.HTML
		Data        any
	}{
		a.Title,
		a.Flash,
		r.URL.Path,
		a.AuthUser,
		csrf.TemplateField(r),
		data,
	}

	ExeTemplate(w, templateFile, res)
}
