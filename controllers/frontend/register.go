package frontend

import (
	"net/http"

	"github.com/faizalom/go-web/config"
	"github.com/faizalom/go-web/lib"

	"github.com/julienschmidt/httprouter"
)

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Template.ParseFiles(config.ThemePath + "/views/register.html")
	Template.ExecuteTemplate(w, "register.html", H{"title": "Register", "google_register_url": lib.GoogleRegisterURL})
}

func CompleteRegister(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Template.ParseFiles(config.ThemePath + "/views/registered.html")
	Template.ExecuteTemplate(w, "registered.html", H{"title": "Register"})
}
