package frontend

import (
	"net/http"

	"github.com/faizalom/go-web/lib"

	"github.com/julienschmidt/httprouter"
)

func Dashboard(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	lib.ExeTemplate(w, "/views/dashboard.html", nil)
}
