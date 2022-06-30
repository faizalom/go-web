package controllers

import (
	"fmt"
	"net/http"

	"github.com/faizalom/go-web/config"
	"github.com/faizalom/go-web/lib"
	"github.com/julienschmidt/httprouter"
)

func LoginIndexContoller(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := config.MyTheme
	t.Title = "Login"
	t.ExeTemp(w, r, "resources/views/login.html", nil)
}

func LoginSubmitContoller(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := lib.SetAuth(w, r)
	if err != nil {
		session := lib.Session(r)
		session.AddFlash(fmt.Sprintln(err), "error")
		session.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
