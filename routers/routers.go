// Package routers create your routes
package routers

import (
	"net/http"

	"github.com/faizalom/go-web/config"

	"github.com/julienschmidt/httprouter"
)

var router = httprouter.New()

/*
Web Routes

Here is where you can register web routes for your application.
Add New Routes inside this function.
*/
func SetRoutes() http.Handler {
	router.ServeFiles("/public/*filepath", http.Dir(config.PublicPath))

	// Call func to define your routes
	web()
	apiRoute()

	router.NotFound = http.HandlerFunc(fileNotFoundHandler)
	router.MethodNotAllowed = http.HandlerFunc(methodNotAllowedHandler)

	return CSRFRoute(router)
}
