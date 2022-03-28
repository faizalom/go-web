package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/faizalom/go-web/lib"
	"github.com/julienschmidt/httprouter"
)

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session := lib.Session(r)
	session.Options.MaxAge = -1
	session.Save(r, w)

	flashes := make(map[string]interface{})
	flashes["error"] = lib.Session(r).Flashes("error")

	lib.Theme.ExeTemp(w, r, "views/login.html", flashes)
}

func LoginSubmit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := lib.SetAuth(w, r)
	if err != nil {
		session := lib.Session(r)
		session.AddFlash(fmt.Sprintln(err), "error")
		session.Save(r, w)
		//session.AddFlash("error", fmt.Sprintln(err))
		//session.AddFlash("Username already taken")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	auth, err := lib.Auth(r)
	if err != nil {
		log.Println(err)
	}

	auth.Options.MaxAge = -1
	err = auth.Save(r, w)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func CoreUIHome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	auth, _ := lib.Auth(r)
	user, _ := lib.MDB.AuthUser(auth)

	data := make(map[string]interface{})
	data["user"] = user

	jsonB, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	jsonStr := string(jsonB)

	resp := struct {
		Json string
	}{
		jsonStr,
	}
	lib.Theme.ExeTemp(w, r, "views/react.html", resp)
}
