package frontend

import (
	"net/http"

	"github.com/faizalom/go-web/lib"

	"github.com/julienschmidt/httprouter"
)

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	lib.ExeTemplate(w, "/views/register.html", H{"title": "Register", "google_register_url": "lib.GoogleRegisterURL"})
}

func CompleteRegister(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	lib.ExeTemplate(w, "/views/registered.html", H{"title": "Register"})
}
