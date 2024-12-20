package frontend

import (
	"net/http"

	"github.com/faizalom/go-web/lib"
)

func Register(w http.ResponseWriter, r *http.Request) {
	lib.ExeTemplate(w, "register.html", H{"title": "Register", "google_register_url": "lib.GoogleRegisterURL"})
}

func CompleteRegister(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.PathValue("jwtToken")
	lib.ExeTemplate(w, "registered.html", H{"title": "Register", "jwtToken": jwtToken})
}
