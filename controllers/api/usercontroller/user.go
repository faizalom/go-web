package usercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/faizalom/go-web/lib"
	"github.com/faizalom/go-web/models"
)

func Profile(w http.ResponseWriter, r *http.Request, auth models.User) {
	lib.Success(w, auth)
}

func Update(w http.ResponseWriter, r *http.Request, auth models.User) {
	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		lib.Error(w, http.StatusBadRequest, "Invalid request. Please input valid input")
		return
	}

	if auth.FirstName == "" {
		lib.Error(w, http.StatusBadRequest, "First name is required")
		return
	}

	if auth.LastName == "" {
		lib.Error(w, http.StatusBadRequest, "Last name is required")
		return
	}

	_, err = models.UpdateUser(auth)
	if err != nil {
		lib.Error(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	lib.Success(w, H{"message": "User updated successfully"})
}
