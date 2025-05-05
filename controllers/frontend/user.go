package frontend

import (
	"net/http"

	"github.com/faizalom/go-web/lib"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	app := lib.GetApp(w, r)
	app.Title = "Welcome to Dashboard"
	app.ExeTemp(w, r, "dashboard.html", nil)
}
