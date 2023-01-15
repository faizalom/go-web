package usercontroller

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/faizalom/go-web/config"
	"github.com/faizalom/go-web/lib"
	"github.com/faizalom/go-web/models"

	"github.com/julienschmidt/httprouter"
)

type H map[string]any

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	request := struct {
		Email      string `json:"email"`
		Password   string `json:"password"`
		RePassword string `json:"repassword"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		lib.Error(w, http.StatusBadRequest, "Invalid request. Please input valid input")
		return
	}

	request.Email = strings.ToLower(strings.TrimSpace(request.Email))
	if request.Email == "" {
		lib.Error(w, http.StatusBadRequest, "Email address is required")
		return
	}

	if !lib.IsEmailValid(request.Email) {
		lib.Error(w, http.StatusBadRequest, "Vaild Email address is required")
		return
	}

	if request.Password == "" {
		lib.Error(w, http.StatusBadRequest, "Password is required")
		return
	}

	if request.Password != request.RePassword {
		lib.Error(w, http.StatusBadRequest, "Password and confirm password must be same")
		return
	}

	passwordHash, err := lib.HashPassword(request.Password)
	if err != nil {
		log.Println(err)
		lib.Error(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	ok, err := models.IsEmailExists(request.Email)
	if err != nil {
		lib.Error(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if ok {
		lib.Error(w, http.StatusBadRequest, "Your Email address already registered")
		return
	}

	// Seperate go routine for sending email to aviod waiting to complete send email
	go func() {
		token, err := lib.J.GenerateJWT(H{"email": request.Email, "password": passwordHash})
		if err != nil {
			log.Println(err)
			return
		}
		// To avoid hacker to insert multi records in DB. User needs to confirm his email
		data := map[string]any{"registration_link": config.ServerURL + "/register/" + token}
		lib.Mail(lib.SendRegisterMail(data, request.Email))
	}()
	lib.Success(w, H{"message": "Confirmation link send to " + request.Email + ". Please complete registration by clicking the link"})
}

func CompleteRegister(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jwtToken := ps.ByName("jwtToken")

	claims, err := lib.J.VerifyJWT(jwtToken)
	if err != nil {
		log.Println(err)
		lib.Error(w, http.StatusBadRequest, "Invalid token or expired")
		return
	}

	user := models.User{
		Email:    claims["email"].(string),
		Password: claims["password"].(string),
	}

	ok, err := models.IsEmailExists(user.Email)
	if err != nil {
		lib.Error(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if ok {
		lib.Error(w, http.StatusBadRequest, "Your Email address already registered")
		return
	}

	userID, err := models.InsertUser(user)
	if err != nil {
		log.Println(err)
		lib.Error(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	token, err := models.GenerateAccessToken(userID)
	if err != nil {
		lib.Error(w, http.StatusBadRequest, "Internal Server Error")
		return
	}
	message := struct {
		Token string `json:"token"`
	}{
		token,
	}
	lib.Success(w, message)
}
