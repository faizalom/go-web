package controllers

import (
	"context"
	"fmt"
	"helper/model"
	"log"
	"net/http"
	"text/template"

	"github.com/faizalom/go-web/lib"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	lib.Theme.ExeTemp(w, r, "views/login.html", nil)
}

func LoginSubmit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	session, _ := lib.Auth(r)

	username := r.FormValue("username")
	password := r.FormValue("password")

	user, err := lib.MDB.Login(username, password)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Values["user_id"] = user.ID.Hex()
	//session.Values["user"] = user
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)

	// // Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	// // if err := r.ParseForm(); err != nil {
	// // 	fmt.Fprintf(w, "ParseForm() err: %v", err)
	// // 	return
	// // }
	// username := r.FormValue("username")
	// password := r.FormValue("password")

	// user, err := lib.MDB.Login(username, password)
	// if err != nil {
	// 	http.Redirect(w, r, "/login", http.StatusSeeOther)
	// 	return
	// }

	// session, _ := lib.Store.Get(r, "cookie-name")
	// session.Values["authenticated"] = true
	// session.Values["user"] = user
	// session.Save(r, w)

	// //http.Redirect(w, r, "/", http.StatusSeeOther)
	// // fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	// // fmt.Fprintf(w, "Username = %s\n", username)
	// // fmt.Fprintf(w, "Password = %s\n", password)
	// // fmt.Fprintln(w, session.Values["username"])
}

func Profile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session, _ := lib.Auth(r)
	userId := fmt.Sprintf("%v", session.Values["user_id"])

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Println("Invalid id")
	}

	// find
	user := model.User{}
	lib.MDB.UserModel().FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&user)
	//fmt.Fprintln(w, user)

	t := template.Must(template.ParseGlob("/home/ocs-11/go/src/github.com/faizalom/go-web/views/*.html"))
	t.ParseFiles("profile.html")

	t.ExecuteTemplate(w, "profile.html", nil)

}
