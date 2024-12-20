package lib_test

import (
	"testing"
	"time"

	"github.com/faizalom/go-web/config"
	"github.com/faizalom/go-web/lib"
	"golang.org/x/crypto/bcrypt"
)

func TestGenerateJWT(t *testing.T) {
	j := lib.JwtStruct{
		SecretKey:       config.Cipher,
		SessionLifetime: time.Duration(5),
	}

	_, err := j.GenerateJWT(map[string]any{"email": "test@email.com"})
	if err != nil {
		t.Error("Error in GenerateJWT")
	}
}
func TestHashPassword(t *testing.T) {
	password := "mysecretpassword"
	hashedPassword, err := lib.HashPassword(password)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		t.Errorf("Hashed password does not match the original password: %v", err)
	}
}

func TestVerifyJWT(t *testing.T) {
	j := lib.JwtStruct{
		SecretKey:       config.Cipher,
		SessionLifetime: time.Duration(5),
	}

	tokenString, err := j.GenerateJWT(map[string]any{"email": "test@email.com"})
	if err != nil {
		t.Errorf("Error generating JWT: %v", err)
	}

	claims, err := j.VerifyJWT(tokenString)
	if err != nil {
		t.Errorf("Error verifying JWT: %v", err)
	}

	if email, ok := claims["email"].(string); !ok || email != "test@email.com" {
		t.Errorf("JWT claims do not match expected values")
	}
}
