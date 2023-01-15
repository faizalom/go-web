package frontend

import (
	"html/template"
	"net/http"

	"github.com/faizalom/go-web/config"
	"github.com/faizalom/go-web/lib"

	"github.com/julienschmidt/httprouter"
)

type H map[string]any

var Template = template.Must(template.ParseGlob(config.ThemePath + "/layout/*.html"))

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Template.ParseFiles(config.ThemePath + "/views/login.html")
	Template.ExecuteTemplate(w, "login.html", H{"title": "Login", "google_login_url": lib.GoogleLoginURL})
}

func GoogleLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Template.ParseFiles(config.ThemePath + "/views/google-login.html")
	Template.ExecuteTemplate(w, "google-login.html", H{"title": "Login"})
}
