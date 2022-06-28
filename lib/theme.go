package lib

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
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

func (t ThemeStruct) ExeTemp(w http.ResponseWriter, r *http.Request, templateFile string, data interface{}) {

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
