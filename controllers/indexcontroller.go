package controllers

import (
	"net/http"

	"github.com/faizalom/go-web/config"
	"github.com/julienschmidt/httprouter"
)

func IndexContoller(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := config.MyTheme
	t.Title = "Wow Works"
	t.ExeTemp(w, r, "resources/views/index.html", nil)
}
