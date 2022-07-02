package lib

import (
	"html/template"
	"log"
	"net/http"

	"github.com/faizalom/go-web/config"
	"github.com/faizalom/go-web/model"
)

type AppStruct struct {
	Template *template.Template
	Title    string
	Flash    map[string]interface{}
	AuthUser model.User
}

var App = AppStruct{}

func init() {

	t := template.Must(template.ParseGlob(config.ThemePath + "*.html"))
	t.ParseGlob(config.ThemeView + "*.html")

	App = AppStruct{
		Template: t,
	}
}

func GetApp(w http.ResponseWriter, r *http.Request) AppStruct {
	session := FlashSession(r)
	session.Options.MaxAge = -1
	session.Save(r, w)

	flash := make(map[string]interface{})
	flash["error"] = session.Flashes("error")
	App.Flash = flash

	auth, e := Auth(r)
	if e != nil {
		log.Println(e)
	}

	App.AuthUser, e = MDB.GetAuthUser(auth)
	if e != nil {
		log.Println(e)
	}

	return App
}
