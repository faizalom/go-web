package routers

import (
	"net/http"

	"github.com/faizalom/go-web/lib"
)

// This function is run user enters invalid url
func fileNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	app := lib.GetApp(w, r)
	app.Title = "404 | Not Found"

	data := struct {
		Message string
	}{
		"404 | Not Found",
	}
	app.ExeTemp(w, r, "resources/views/error.html", data)
}

// This function is run user sends a wrong method
func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	app := lib.GetApp(w, r)
	app.Title = "405 | Method Not Allowed"

	data := struct {
		Message string
	}{
		"405 | Method Not Allowed",
	}
	app.ExeTemp(w, r, "resources/views/error.html", data)
}

// This function is run during the CSRF token expired
func pageExpiredHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	app := lib.GetApp(w, r)
	app.Title = "404 | Page Expired"

	data := struct {
		Message string
	}{
		"403 | Page Expired",
	}
	app.ExeTemp(w, r, "resources/views/error.html", data)
}

// func unauthorizedHandler(w http.ResponseWriter, r *http.Request) {
// 	http.Error(w, fmt.Sprintf("%s - %s", http.StatusText(http.StatusForbidden), "ewrew"), http.StatusForbidden)
// }
