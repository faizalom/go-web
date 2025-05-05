package lib

import (
	"log"
	"net/http"

	"github.com/faizalom/go-web/config"
	"github.com/faizalom/go-web/models"
	"github.com/gorilla/sessions"
)

var (
	// Sessionkey is the key used to encrypt the session data
	// sessionKey will be assigned in the config file (application.yaml) with the key "cipherkey"
	Sessionkey []byte
	Store      *sessions.CookieStore
)

func InitSession() {
	Sessionkey = []byte(config.Cipher)
	Store = sessions.NewCookieStore(Sessionkey)
}

func GetSession(r *http.Request) (*sessions.Session, error) {
	return Store.Get(r, "session")
}

func AddSession(w http.ResponseWriter, r *http.Request, key string, value any) {
	session, e := GetSession(r)
	if e != nil {
		log.Println(e)
	}

	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   config.SessionLifetime * 60,
		Secure:   true,
		HttpOnly: true, // no websocket or any protocol else
	}

	session.Values[key] = value
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func FlashSession(r *http.Request) *sessions.Session {
	session, e := Store.Get(r, "flash")
	if e != nil {
		log.Println(e)
	}
	return session
}

func AddFlash(w http.ResponseWriter, r *http.Request, message string, key string) {
	session := FlashSession(r)
	session.AddFlash(message, key)

	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetFlashes(w http.ResponseWriter, r *http.Request, key string) []interface{} {
	session := FlashSession(r)
	flashes := session.Flashes(key)

	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	return flashes
}

func Auth(r *http.Request) (*sessions.Session, error) {
	return Store.Get(r, "auth")
}

func SetAuth(w http.ResponseWriter, r *http.Request) string {
	auth, e := Auth(r)
	if e != nil {
		log.Println(e)
	}
	auth.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   config.SessionLifetime * 60,
		Secure:   true,
		HttpOnly: true,
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := models.Login(email, password)
	if err != "" {
		return err
	}

	// Set user as authenticated
	auth.Values["authenticated"] = true
	auth.Values["user_id"] = user.ID
	auth.Save(r, w)
	return ""
}

func SetAuthByID(w http.ResponseWriter, r *http.Request, userID int64) string {
	auth, e := Auth(r)
	if e != nil {
		log.Println(e)
	}
	auth.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   config.SessionLifetime * 60,
		Secure:   true,
		HttpOnly: true,
	}

	// Set user as authenticated
	auth.Values["authenticated"] = true
	auth.Values["user_id"] = userID
	auth.Save(r, w)
	return ""
}

func ClearAuth(w http.ResponseWriter, r *http.Request) {
	session, e := Auth(r)
	if e != nil {
		log.Println(e)
	}
	session.Options.MaxAge = -1
	session.Save(r, w)

	session = FlashSession(r)
	if e != nil {
		log.Println(e)
	}
	session.Options.MaxAge = -1
	session.Save(r, w)

	session, e = GetSession(r)
	if e != nil {
		log.Println(e)
	}
	session.Options.MaxAge = -1
	session.Save(r, w)
}
