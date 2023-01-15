package routers

import (
	"net/http"

	"html/template"

	"github.com/faizalom/go-web/config"
)

var Template = template.Must(template.ParseGlob(config.ThemePath + "/layout/*.html"))

type H map[string]any

// This function is run user enters invalid url
func fileNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	Template.ParseFiles(config.ThemePath + "/views/error.html")
	Template.ExecuteTemplate(w, "error.html", H{"title": "404 | Not Found"})
}

// This function is run user sends a wrong method
func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	Template.ParseFiles(config.ThemePath + "/views/error.html")
	Template.ExecuteTemplate(w, "error.html", H{"title": "405 | Method Not Allowed"})
}
