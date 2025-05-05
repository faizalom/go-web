package frontend

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/faizalom/go-web/lib"
	"github.com/faizalom/go-web/models"
)

type H map[string]any

func Login(w http.ResponseWriter, r *http.Request) {
	googleUrl, state := lib.GetGoogleLoginURL()
	lib.AddSession(w, r, "google_login_state", state)

	app := lib.GetApp(w, r)
	app.Title = "Login"
	app.ExeTemp(w, r, "login.html", H{"google_login_url": googleUrl})
}

func LoginSubmit(w http.ResponseWriter, r *http.Request) {
	errMessage := lib.SetAuth(w, r)
	if errMessage != "" {
		lib.RedirectWithError(w, r, errMessage, "/login")
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	app := lib.GetApp(w, r)
	app.Title = "Google Login"

	session, err := lib.GetSession(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		app.ExeTemp(w, r, "message.html", H{"message": strconv.Itoa(http.StatusInternalServerError) + " | " + "Internal Server Error"})
		return
	}
	sessionLoginState, okLogin := session.Values["google_login_state"].(string)
	sessionRegState, okReg := session.Values["google_register_state"].(string)
	if !okLogin && !okReg {
		w.WriteHeader(http.StatusBadRequest)
		app.ExeTemp(w, r, "message.html", H{"message": strconv.Itoa(http.StatusBadRequest) + " | " + "Invalid session state"})
	}

	state := r.URL.Query().Get("state")
	code := r.URL.Query().Get("code")
	user, err := lib.GoogleGetUserInfo(code)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		app.ExeTemp(w, r, "message.html", H{"message": strconv.Itoa(http.StatusBadRequest) + " | " + "Authentication Failed"})
		return
	}

	if state == sessionLoginState {
		// Check if the user exists in the database
		user, err = models.GetUserByGoogleId(user.GoogleID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			app.ExeTemp(w, r, "message.html", H{"message": strconv.Itoa(http.StatusInternalServerError) + " | " + "Internal Server Error"})
			return
		}
		userID := user.ID

		if userID == 0 {
			app.ExeTemp(w, r, "message.html", H{"message": template.HTML("This google accout is not matched with our records. <br />Please <a href='/register'>register</a> with us and continue")})
			return
		}

		lib.SetAuthByID(w, r, userID)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if state == sessionRegState {
		// Insert the user into the database
		userID, err := models.InsertUserWithDetails(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			app.ExeTemp(w, r, "message.html", H{"message": strconv.Itoa(http.StatusInternalServerError) + " | " + "Internal Server Error"})
			return
		}

		lib.SetAuthByID(w, r, userID)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	lib.ClearAuth(w, r)
	http.Redirect(w, r, "/login", http.StatusFound)
}
