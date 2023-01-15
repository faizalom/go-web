package lib_test

import (
	"testing"
	"time"

	"github.com/faizalom/go-web/config"
	"github.com/faizalom/go-web/lib"
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
