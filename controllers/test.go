package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/faizalom/go-web/lib"
	"github.com/julienschmidt/httprouter"
)

func CoreUI(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	auth, _ := lib.Auth(r)
	user, _ := lib.MDB.AuthUser(auth)

	data := make(map[string]interface{})
	data["user"] = user

	jsonB, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	jsonStr := string(jsonB)

	resp := struct {
		Json string
	}{
		jsonStr,
	}
	lib.Theme.ExeTemp(w, r, "views/react.html", resp)
}
