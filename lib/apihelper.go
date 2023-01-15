package lib

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/faizalom/go-web/config"
)

var J JwtStruct

func init() {
	J.SecretKey = config.Cipher
	J.SessionLifetime = time.Duration(5)
}

func Success(w http.ResponseWriter, data any) {
	payload, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}

func Error(w http.ResponseWriter, statusCode int, message string) {
	errMessage := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		"error",
		message,
	}

	payload, err := json.Marshal(errMessage)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(payload))
}
