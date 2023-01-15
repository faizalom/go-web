package frontend

import (
	"net/http"

	"github.com/faizalom/go-web/config"

	"github.com/julienschmidt/httprouter"
)

func Dashboard(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Template.ParseFiles(config.ThemePath + "/views/dashboard.html")
	Template.ExecuteTemplate(w, "dashboard.html", nil)
}
