package lib

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/faizalom/go-web/model"
	"github.com/gorilla/csrf"
)

// function to run html template
func (a AppStruct) ExeTemp(w http.ResponseWriter, r *http.Request, templateFile string, data interface{}) {
	var res = struct {
		Title       string
		Flash       map[string]interface{}
		CurrentPath string
		AuthUser    model.User
		CsrfField   template.HTML
		Data        interface{}
	}{
		a.Title,
		a.Flash,
		r.URL.Path,
		a.AuthUser,
		csrf.TemplateField(r),
		data,
	}

	a.Template.ParseFiles(templateFile)
	s := strings.Split(templateFile, "/")
	templateName := s[len(s)-1]
	a.Template.ExecuteTemplate(w, templateName, res)
}
