package apis

import (
	"context"
	"encoding/json"
	"helper/model"
	"log"
	"net/http"

	"github.com/faizalom/go-web/lib"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetStaff(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users := []model.User{}
	cur, err := lib.MDB.UserModel().Find(context.Background(), bson.D{})
	if err != nil {
		log.Println(err)
	}
	defer cur.Close(context.Background())
	if err = cur.All(context.Background(), &users); err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Println(err)
	}
}

func GetStaffByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	objID, err := primitive.ObjectIDFromHex(p.ByName("id"))
	if err != nil {
		log.Println(err)
	}

	user := model.User{}
	err = lib.MDB.UserModel().FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println(err)
	}
}

func CheckStaffAvailable(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	data := make(map[string]string)
	err := decoder.Decode(&data)
	if err != nil {
		log.Println(err)
	}

	objID, err := primitive.ObjectIDFromHex(data["id"])
	if err != nil {
		log.Println(err)
	}

	user := model.User{}
	err = lib.MDB.UserModel().FindOne(context.Background(), bson.M{
		"_id":      bson.M{"$ne": objID},
		"username": data["username"],
	}).Decode(&user)
	if err != nil {
		log.Println(err)
	}

	res := struct {
		IsAvailable bool `json:"is_available"`
	}{}
	if user.Username == "" {
		res.IsAvailable = true
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println(err)
	}
}

func UpdateStaffByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	objID, err := primitive.ObjectIDFromHex(p.ByName("id"))
	if err != nil {
		log.Println(err)
	}

	decoder := json.NewDecoder(r.Body)
	user := model.User{}
	err = decoder.Decode(&user)
	if err != nil {
		log.Println(err)
	}

	_, err = lib.MDB.UserModel().ReplaceOne(context.Background(), bson.M{"_id": objID}, user)
	if err != nil {
		log.Println(err)
	}

	res := struct {
		Message string `json:"message"`
	}{
		"Updated Successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println(err)
	}
}

func StaffStore(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	decoder := json.NewDecoder(r.Body)
	user := model.User{}
	err := decoder.Decode(&user)
	if err != nil {
		log.Println(err)
	}

	_, err = lib.MDB.UserModel().InsertOne(context.Background(), user)
	if err != nil {
		log.Println(err)
	}

	res := struct {
		Message string `json:"message"`
	}{
		"User added successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println(err)
	}
}
