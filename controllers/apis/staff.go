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
