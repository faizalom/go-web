package frontend

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/faizalom/go-web/config"
	"github.com/faizalom/go-web/lib"
	"github.com/faizalom/go-web/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	googleUrl, state := lib.GetGoogleLoginURL()
	lib.AddSession(w, r, "google_register_state", state)

	app := lib.GetApp(w, r)
	app.Title = "Register"
	app.ExeTemp(w, r, "register.html", H{"google_register_url": googleUrl})
}

func RegisterSubmit(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	repassword := r.FormValue("repassword")

	email = strings.ToLower(strings.TrimSpace(email))
	if email == "" {
		lib.RedirectWithError(w, r, "Email address is required", "/register")
		return
	}

	if !lib.IsEmailValid(email) {
		lib.RedirectWithError(w, r, "Vaild Email address is required", "/register")
		return
	}

	if password == "" {
		lib.RedirectWithError(w, r, "Password is required", "/register")
		return
	}

	if password != repassword {
		lib.RedirectWithError(w, r, "Password and confirm password must be same", "/register")
		return
	}

	passwordHash, err := lib.HashPassword(password)
	if err != nil {
		log.Println(err)
		lib.RedirectWithError(w, r, "Internal Server Error", "/register")
		return
	}

	ok, err := models.IsEmailExists(email)
	if err != nil {
		lib.RedirectWithError(w, r, "Internal Server Error", "/register")
		return
	}

	if ok {
		lib.RedirectWithError(w, r, "Your Email address already registered", "/register")
		return
	}

	// Seperate go routine for sending email to aviod waiting to complete send email
	go func() {
		token, err := lib.J.GenerateJWT(H{"email": email, "password": passwordHash})
		if err != nil {
			log.Println(err)
			return
		}
		// To avoid hacker to insert multi records in DB. User needs to confirm his email
		data := map[string]any{"registration_link": config.Server.URL + "/register/" + token}
		lib.Mail(lib.SendRegisterMail(data, email))
	}()
	// WIP
	app := lib.GetApp(w, r)
	app.Title = "Register"
	app.ExeTemp(w, r, "message.html", H{
		"message": template.HTML("Confirmation link send to " + email + " <br />Please complete registration by clicking the registration link"),
		"type":    "success",
	})
}

func CompleteRegister(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.PathValue("jwtToken")

	app := lib.GetApp(w, r)
	app.Title = "Register"

	claims, err := lib.J.VerifyJWT(jwtToken)
	if err != nil {
		log.Println(err)
		app.ExeTemp(w, r, "message.html", H{"message": "Invalid token or expired"})
		return
	}

	user := models.User{
		Email:    claims["email"].(string),
		Password: claims["password"].(string),
	}

	ok, err := models.IsEmailExists(user.Email)
	if err != nil {
		app.ExeTemp(w, r, "message.html", H{"message": "Internal Server Error"})
		return
	}
	if ok {
		app.ExeTemp(w, r, "message.html", H{"message": "Your Email address already registered"})
		return
	}

	_, err = models.InsertUser(user)
	if err != nil {
		log.Println(err)
		app.ExeTemp(w, r, "message.html", H{"message": "Internal Server Error"})
		return
	}

	app.ExeTemp(w, r, "message.html", H{
		"message": template.HTML("Registration successful! <br />You can now <a href='/login'>Log In</a> with your credentials."),
		"type":    "success",
	})
}
