package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/faizalom/go-web/lib"
	"github.com/faizalom/go-web/model"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LoginIndexController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	app := lib.GetApp(w, r)
	app.Title = "Login"
	app.ExeTemp(w, r, "resources/views/login.html", nil)
}

func LoginSubmitController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := lib.SetAuth(w, r)
	if err != nil {
		session := lib.FlashSession(r)
		session.AddFlash(fmt.Sprintln(err), "error")
		session.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func LogoutController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session := lib.FlashSession(r)
	session.Options.MaxAge = -1
	session.Save(r, w)

	auth, e := lib.Auth(r)
	if e != nil {
		log.Println(e)
	}
	auth.Options.MaxAge = -1
	auth.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func RegisterController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	app := lib.GetApp(w, r)
	app.Title = "Register"
	app.ExeTemp(w, r, "resources/views/register.html", nil)
}

func RegisterSubmitController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	isValid := Validate(w, r)
	if !isValid {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	user := model.User{}
	user.FirstName = r.FormValue("first_name")
	user.LastName = r.FormValue("last_name")
	user.Email = r.FormValue("email")
	password, e := lib.HashPassword(r.FormValue("password"))
	if e != nil {
		log.Println(e)
	}
	user.Password = password

	opts := options.InsertOne()
	lib.MDB.UserModel().InsertOne(context.TODO(), user, opts)

	lib.SetAuth(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Validate(w http.ResponseWriter, r *http.Request) bool {
	isValid := true
	session := lib.FlashSession(r)

	var rxEmail = regexp.MustCompile(".+@.+\\..+")
	match := rxEmail.Match([]byte(r.FormValue("email")))
	if match == false {
		session.AddFlash("Please enter a valid email address", "error")
	}

	if strings.TrimSpace(r.FormValue("first_name")) == "" {
		session.AddFlash("Please enter a first name", "error")
	}

	if strings.TrimSpace(r.FormValue("last_name")) == "" {
		session.AddFlash("Please enter a last name", "error")
	}
	session.Save(r, w)

	if len(session.Flashes("error")) > 0 {
		isValid = false
	}
	return isValid
}
