package routers

import (
	"net/http"

	"github.com/faizalom/go-web/config"
	"github.com/faizalom/go-web/lib"
)

type H map[string]any

// This function is run user enters invalid url
func fileNotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	lib.Template.ParseFiles(config.Path.Theme + "/views/error.html")
	lib.Template.ExecuteTemplate(w, "error.html", H{"title": "404 | Not Found"})
}

// This function is run user sends a wrong method
func methodNotAllowedHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	lib.Template.ParseFiles(config.Path.Theme + "/views/error.html")
	lib.Template.ExecuteTemplate(w, "error.html", H{"title": "405 | Method Not Allowed"})
}

// This function is run during the CSRF token expired
func pageExpiredHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(419)
	lib.Template.ParseFiles(config.Path.Theme + "/views/error.html")
	lib.Template.ExecuteTemplate(w, "error.html", H{"title": "419 | Page Expired"})
}
