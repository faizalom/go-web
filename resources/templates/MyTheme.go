package templates

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Navlink struct {
	Link     string
	Icon     string
	Text     string
	Children []Navlink
}

type ThemeStruct struct {
	Template *template.Template
	Title    string
	SideMenu []Navlink
}

//var Path, _ = os.Getwd()
var ThemePath = "resources/templates/"

//var ThemePath = "/home/ocs-11/Dropbox/go/src/theme/CoreUI/"
var MyTheme ThemeStruct

func init() {
	MyTheme.Template = template.Must(template.ParseGlob(ThemePath + "*.html"))
}

func (t ThemeStruct) ExeTemp(w http.ResponseWriter, r *http.Request, templateFile string, data any) {

	jsonB, err := json.Marshal(t.SideMenu)
	if err != nil {
		log.Println(err)
	}
	navlinkJson := string(jsonB)

	var res = struct {
		Title       string
		Menu        []Navlink
		MenuJson    string
		CurrentPath string
		Data        interface{}
	}{
		t.Title,
		t.SideMenu,
		navlinkJson,
		r.URL.Path,
		data,
	}

	t.Template.ParseFiles(templateFile)
	s := strings.Split(templateFile, "/")
	templateName := s[len(s)-1]
	t.Template.ExecuteTemplate(w, templateName, res)
}
