package lib

import (
	"net/http"
	"strings"

	"github.com/faizalom/go-web/model"
)

func (a AppStruct) ExeTemp(w http.ResponseWriter, r *http.Request, templateFile string, data interface{}) {

	var res = struct {
		Title       string
		Flash       map[string]interface{}
		CurrentPath string
		AuthUser    model.User
		Data        interface{}
	}{
		a.Title,
		a.Flash,
		r.URL.Path,
		a.AuthUser,
		data,
	}

	a.Template.ParseFiles(templateFile)
	s := strings.Split(templateFile, "/")
	templateName := s[len(s)-1]
	a.Template.ExecuteTemplate(w, templateName, res)
}
