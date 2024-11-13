package frontend

import (
	"net/http"

	"github.com/faizalom/go-web/lib"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	lib.ExeTemplate(w, "profile.html", H{"title": "Profile"})
}
