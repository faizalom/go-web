package lib

import (
	"encoding/json"
	"log"
	"net/http"
)

// REMOVE

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
		Message string `json:"message"`
		Status  string `json:"status"`
	}{
		message,
		"error",
	}

	payload, err := json.Marshal(errMessage)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(payload))
}
