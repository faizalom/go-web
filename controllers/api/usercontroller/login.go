package usercontroller

import (
	"database/sql"
	"net/http"

	"github.com/faizalom/go-web/lib"
	"github.com/faizalom/go-web/models"

	"github.com/julienschmidt/httprouter"
)

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	email, password, ok := r.BasicAuth()
	if !ok {
		lib.Error(w, http.StatusBadRequest, "Invalid Credentials")
		return
	}
	userID, err := models.Login(email, password)
	if err == sql.ErrNoRows {
		lib.Error(w, http.StatusForbidden, "Username entered does not exist")
		return
	}

	if err != nil {
		lib.Error(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if userID == 0 {
		lib.Error(w, http.StatusForbidden, "Password is incorrect")
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

func GoogleLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	/*
		state := r.URL.Query().Get("state")
		code := r.URL.Query().Get("code")
		user, err := lib.GoogleGetUserInfo(code)
		if err != nil {
			lib.Error(w, http.StatusBadRequest, "Authentication Failed")
			return
		}

		var userID int64
		redirectTo := config.ServerURL

		if state == lib.RegisterState {
			ok, err := models.IsEmailExists(user.Email)
			if err != nil {
				lib.Error(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			if ok {
				lib.Error(w, http.StatusBadRequest, "Your Email address already registered")
				return
			}

			userID, err = models.InsertUserWithDetails(user)
			if err != nil {
				lib.Error(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
			redirectTo = config.ServerURL + "/profile"
		} else {
			user, err = models.GetUserByGoogleId(user.GoogleID)
			if err != nil {
				lib.Error(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
			userID = user.ID

			if userID == 0 {
				lib.Error(w, http.StatusForbidden, "This google accout is not matched with our records. Please register with us and continue")
				return
			}
		}

		token, err := models.GenerateAccessToken(userID)
		if err != nil {
			lib.Error(w, http.StatusBadRequest, "Internal Server Error")
			return
		}
		message := struct {
			Token      string `json:"token"`
			RedirectTo string `json:"redirect_to"`
		}{
			token,
			redirectTo,
		}
		lib.Success(w, message)
	*/
}
