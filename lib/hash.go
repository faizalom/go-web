package lib

import (
	"fmt"
	"time"

	"maps"

	"github.com/faizalom/go-web/config"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type JwtStruct struct {
	SecretKey       string
	SessionLifetime time.Duration
}

var J JwtStruct

func InitHash() {
	J.SecretKey = config.Cipher
	J.SessionLifetime = time.Duration(5)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Generate JWT token
func (j JwtStruct) GenerateJWT(cli map[string]any) (string, error) {
	var mySigningKey = []byte(j.SecretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	//claims["authorized"] = true
	maps.Copy(claims, cli)
	claims["exp"] = time.Now().Add(time.Minute * j.SessionLifetime).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j JwtStruct) VerifyJWT(tokenString string) (jwt.MapClaims, error) {
	var mySigningKey = []byte(j.SecretKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing token")
		}
		return mySigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("there was an error in parsing token")
}
