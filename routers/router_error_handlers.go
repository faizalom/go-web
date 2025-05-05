package routers

import (
	"net/http"

	"github.com/faizalom/go-web/lib"
)

type H map[string]any

// This function is run user enters invalid url
func fileNotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	lib.Template.ExecuteTemplate(w, "message.html", H{"Title": "404 | Page not found", "Data": map[string]string{"message": "404 | Page not found"}})
}

// This function is run during the CSRF token expired
func pageExpiredHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(419)
	lib.Template.ExecuteTemplate(w, "message.html", H{"Title": "419 | Page Expired", "Data": map[string]string{"message": "419 | Page Expired"}})
}
