package controllers

import (
	"fmt"
	"net/http"

	"github.com/faizalom/go-web/lib"
	"github.com/julienschmidt/httprouter"
)

func LoginIndexController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	app := lib.GetApp(w, r)
	app.Title = "Login"
	app.ExeTemp(w, r, "resources/views/login.html", nil)
}

func LoginSubmitController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := lib.SetAuth(w, r)
	if err != nil {
		session := lib.FlashSession(r)
		session.AddFlash(fmt.Sprintln(err), "error")
		session.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
