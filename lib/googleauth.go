package lib

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/faizalom/go-web/config"
	"github.com/faizalom/go-web/models"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetGoogleLoginURL() (string, string) {
	var googleOAuthConfig = &oauth2.Config{
		ClientID:     config.GoogleKey.ClientID,
		ClientSecret: config.GoogleKey.ClientSecret,
		RedirectURL:  config.Server.URL + config.GoogleKey.CallbackURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		},
		Endpoint: google.Endpoint,
	}

	state := RandToken()
	return googleOAuthConfig.AuthCodeURL(state), state
}

func GoogleGetUserInfo(code string) (models.User, error) {
	var googleOAuthConfig = &oauth2.Config{
		ClientID:     config.GoogleKey.ClientID,
		ClientSecret: config.GoogleKey.ClientSecret,
		RedirectURL:  config.Server.URL + config.GoogleKey.CallbackURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		},
		Endpoint: google.Endpoint,
	}

	var user models.User

	ctx, cancelfunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelfunc()

	tok, err := googleOAuthConfig.Exchange(ctx, code)
	if err != nil {
		log.Println(err)
		return user, err
	}

	client := googleOAuthConfig.Client(ctx, tok)
	res, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		log.Println(err)
		return user, err
	}
	defer res.Body.Close()

	var userDetails struct {
		GoogleID  string `json:"sub"`
		FirstName string `json:"given_name"`
		LastName  string `json:"family_name"`
		Email     string `json:"email"`
	}

	err = json.NewDecoder(res.Body).Decode(&userDetails)
	if err != nil {
		log.Println(err)
		return user, err
	}

	user.GoogleID = userDetails.GoogleID
	user.FirstName = userDetails.FirstName
	user.LastName = userDetails.LastName
	user.Email = userDetails.Email

	return user, err
}
