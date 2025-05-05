// Package Middleware handles an incoming request
// Users can create new Middleware functions here or create a new file beneath the middleware directory
package middleware

import (
	"net/http"

	"github.com/faizalom/go-web/lib"
)

// Users should be redirected to log in when they are not authenticated
func AuthMiddleware(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		auth, err := lib.Auth(r)
		if auth, ok := auth.Values["authenticated"].(bool); !ok || !auth || err != nil {
			redirectToLogin(w, r)
		} else {
			// Original function call
			f(w, r)
		}
	}
}

// This function is called from AuthMiddleware
// This will run at user tries to access auth routes with invalid or Expired auth session or without login
// User can use this fucntion in controllers
func redirectToLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusFound)
}
