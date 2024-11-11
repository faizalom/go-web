package lib

import (
	"net/http"
	"text/template"

	"github.com/faizalom/go-web/config"
)

// func init() {
// 	Template.ParseGlob("D:/Users/Faizal/Dropbox/go/src/projects/coindcx-api/template/*.html")
// }

// function to run html template
// func ExeTemplate(w http.ResponseWriter, templateFile string, data any) {
// 	Template.ParseFiles(config.ThemePath + templateFile)
// 	s := strings.Split(templateFile, "/")
// 	templateName := s[len(s)-1]
// 	err := Template.ExecuteTemplate(w, templateName, data)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func ExeTemplate(w http.ResponseWriter, templateFile string, data any) {
	var Template = template.Must(template.ParseGlob(config.ThemePath + "/*.html"))

	err := Template.ExecuteTemplate(w, templateFile, data)
	if err != nil {
		panic(err)
	}
}
