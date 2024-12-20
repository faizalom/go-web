package frontend

import (
	"net/http"

	"github.com/faizalom/go-web/lib"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	lib.ExeTemplate(w, "dashboard.html", H{"title": "Welcome"})
}
