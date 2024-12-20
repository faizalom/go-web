package lib_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/faizalom/go-web/lib"
)

func TestSuccess(t *testing.T) {
	rr := httptest.NewRecorder()
	data := map[string]string{"message": "success"}

	lib.Success(rr, data)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected, _ := json.Marshal(data)
	if rr.Body.String() != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), string(expected))
	}
}

func TestError(t *testing.T) {
	rr := httptest.NewRecorder()
	statusCode := http.StatusBadRequest
	message := "bad request"

	lib.Error(rr, statusCode, message)

	if status := rr.Code; status != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", status, statusCode)
	}

	expected, _ := json.Marshal(map[string]string{"status": "error", "message": message})
	if rr.Body.String() != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), string(expected))
	}
}
