package routers

import (
	"html/template"
	"net/http"

	"github.com/faizalom/go-web/config"
)

var Template = template.Must(template.ParseGlob(config.ThemePath))

type H map[string]any

// This function is run user enters invalid url
func fileNotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	Template.ParseFiles(config.ThemePath + "/views/error.html")
	Template.ExecuteTemplate(w, "error.html", H{"title": "404 | Not Found"})
}

// This function is run user sends a wrong method
func methodNotAllowedHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	Template.ParseFiles(config.ThemePath + "/views/error.html")
	Template.ExecuteTemplate(w, "error.html", H{"title": "405 | Method Not Allowed"})
}

// This function is run during the CSRF token expired
func pageExpiredHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(419)
	Template.ParseFiles(config.ThemePath + "/views/error.html")
	Template.ExecuteTemplate(w, "error.html", H{"title": "419 | Page Expired"})
}
