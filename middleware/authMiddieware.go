// Package Middleware handle an incoming request
package middleware

import (
	"net/http"

	"github.com/faizalom/go-web/lib"
	"github.com/julienschmidt/httprouter"
)

// User should be redirected to login when they are not authenticated.
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
