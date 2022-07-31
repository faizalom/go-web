// Package Middleware handles an incoming request
// Users can create new Middleware functions here or create a new file beneath the middleware directory
package middleware

import (
	"net/http"

	"github.com/faizalom/go-web/lib"
	"github.com/julienschmidt/httprouter"
)

// Users should be redirected to log in when they are not authenticated
func AuthMiddleware(f func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		auth, _ := lib.Auth(r)
		if auth, ok := auth.Values["authenticated"].(bool); !ok || !auth {
			redirectToLogin(w, r, ps)
		} else {
			// Original function call
			f(w, r, ps)
		}
	}
}
