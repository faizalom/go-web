package lib

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	Store = sessions.NewCookieStore(key)
	//var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
)

func Auth(r *http.Request) (*sessions.Session, error) {
	return Store.Get(r, "auth")
}

func SetAuth(w http.ResponseWriter, r *http.Request) error {
	auth, _ := Auth(r)

	username := r.FormValue("username")
	password := r.FormValue("password")

	user, err := MDB.Login(username, password)
	if err != nil {
		return err
	}

	// Set user as authenticated
	auth.Values["authenticated"] = true
	auth.Values["user_id"] = user.ID.Hex()
	auth.Save(r, w)

	return nil
}

func Session(r *http.Request) *sessions.Session {
	session, e := Store.Get(r, "session")
	if e != nil {
		log.Println(e)
	}
	return session
}
