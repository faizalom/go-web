package frontend

import (
	"net/http"

	"github.com/faizalom/go-web/lib"

	"github.com/julienschmidt/httprouter"
)

func Profile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	lib.ExeTemplate(w, "/views/profile.html", nil)
}
