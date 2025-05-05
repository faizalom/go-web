package lib

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"regexp"
)

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func RedirectWithError(w http.ResponseWriter, r *http.Request, errMessage string, redirect string) {
	AddFlash(w, r, errMessage, "error")
	http.Redirect(w, r, redirect, http.StatusFound)
}
