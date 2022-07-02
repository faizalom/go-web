package controllers

import (
	"net/http"

	"github.com/faizalom/go-web/lib"
	"github.com/julienschmidt/httprouter"
)

func IndexController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	app := lib.GetApp(w, r)
	app.Title = "Dashboard"
	app.ExeTemp(w, r, "resources/views/index.html", nil)
}
