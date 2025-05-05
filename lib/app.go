// Package Libraries
package lib

import (
	"log"
	"net/http"

	"github.com/faizalom/go-web/models"
)

type AppStruct struct {
	Title    string
	Flash    map[string]interface{}
	AuthUser models.User
}

var App = AppStruct{}

func GetApp(w http.ResponseWriter, r *http.Request) AppStruct {
	flash := make(map[string]any)
	flash["error"] = GetFlashes(w, r, "error")
	flash["success"] = GetFlashes(w, r, "success")
	App.Flash = flash

	auth, e := Auth(r)
	if e != nil {
		log.Println(e)
	}

	App.AuthUser, e = models.GetAuthUser(auth)
	if e != nil {
		log.Println(e)
	}

	return App
}
