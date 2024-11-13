// Package Middleware handles an incoming request
// Users can create new Middleware functions here or create a new file beneath the middleware directory
package middleware

import (
	"net/http"
	"strings"

	"github.com/faizalom/go-web/lib"
	"github.com/faizalom/go-web/models"
)

// Users should be redirected to log in when they are not authenticated
func ApiAuthMiddleware(f func(http.ResponseWriter, *http.Request, models.User)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			lib.Error(w, http.StatusUnauthorized, "Token not available")
			return
		}

		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) == 2 {
			reqToken = splitToken[1]
		} else {
			lib.Error(w, http.StatusUnauthorized, "Token not available")
			return
		}

		if r.URL.Path == "/api/logout" {
			models.DeleteAccessToken(reqToken)
			message := struct {
				Token string `json:"token"`
			}{
				"",
			}
			lib.Success(w, message)
			return
		}

		if auth, err := models.GetAuthUser(reqToken); err == nil {
			if auth.ID == 0 {
				lib.Error(w, http.StatusUnauthorized, "User not found")
			} else {
				f(w, r, auth)
			}
		} else {
			lib.Error(w, http.StatusInternalServerError, "Internal Server Error")
		}
	}
}
