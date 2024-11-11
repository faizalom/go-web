package frontend

import (
	"net/http"

	"github.com/faizalom/go-web/lib"

	"github.com/julienschmidt/httprouter"
)

type H map[string]any

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	lib.ExeTemplate(w, "/views/login.html", H{"title": "Login", "google_login_url": "lib.GoogleLoginURL"})
}

func GoogleLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	lib.ExeTemplate(w, "/views/google-login.html", H{"title": "Login"})
}
