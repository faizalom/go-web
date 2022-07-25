// Package Middleware
package middleware

import (
	"net/http"

	"github.com/faizalom/go-web/lib"
	"github.com/julienschmidt/httprouter"
)

func AuthMiddleware(f func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		auth, _ := lib.Auth(r)
		if auth, ok := auth.Values["authenticated"].(bool); !ok || !auth {
			redirectToLogin(w, r, ps)
		} else {
			f(w, r, ps) // original function call
		}
	}
}
