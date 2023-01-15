package frontend

import (
	"net/http"

	"github.com/faizalom/go-web/config"

	"github.com/julienschmidt/httprouter"
)

func Profile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Template.ParseFiles(config.ThemePath + "/views/profile.html")
	Template.ExecuteTemplate(w, "profile.html", nil)
}
