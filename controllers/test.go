package controllers

import (
	"net/http"

	"github.com/faizalom/go-web/lib"
	"github.com/julienschmidt/httprouter"
)

func CoreUI(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	lib.Theme.ExeTemp(w, r, "views/dashboard.html", nil)
}
