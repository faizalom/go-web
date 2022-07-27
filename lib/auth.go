package lib

import (
	"log"
	"net/http"

	"github.com/faizalom/go-web/config"
	"github.com/gorilla/sessions"
)

var (
	key   = []byte(config.Cipher)
	Store = sessions.NewCookieStore(key)
)

func Auth(r *http.Request) (*sessions.Session, error) {
	return Store.Get(r, "auth")
}

func SetAuth(w http.ResponseWriter, r *http.Request) error {
	auth, e := Auth(r)
	if e != nil {
		log.Println(e)
	}
	auth.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   config.SessionLifetime * 60,
		Secure:   true,
		HttpOnly: true, // no websocket or any protocol else
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := MDB.Login(email, password)
	if err != nil {
		return err
	}

	// Set user as authenticated
	auth.Values["authenticated"] = true
	auth.Values["user_id"] = user.ID.Hex()
	auth.Save(r, w)
	return nil
}

func FlashSession(r *http.Request) *sessions.Session {
	session, e := Store.Get(r, "flash")
	if e != nil {
		log.Println(e)
	}
	return session
}
