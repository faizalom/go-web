package model

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	MemberCode   string             `json:"memberCode" bson:"memberCode"`
	FirstName    string             `json:"firstName" bson:"firstName"`
	LastName     string             `json:"lastName" bson:"lastName"`
	Email        string             `json:"email" bson:"email"`
	AllowLogin   bool               `json:"allow_login" bson:"allow_login"`
	Username     string             `json:"username" bson:"username"`
	Password     string             `json:"password" bson:"password"`
	Mobile       string             `json:"mobile" bson:"mobile"`
	Address1     string             `json:"address_1" bson:"address_1"`
	Address2     string             `json:"address_2" bson:"address_2"`
	City         string             `json:"city" bson:"city"`
	PostalCode   string             `json:"postal_code" bson:"postal_code"`
	DOB          string             `json:"dob" bson:"dob"`
	Comments     string             `json:"comments" bson:"comments"`
	BloodGroup   string             `json:"blood_group" bson:"blood_group"`
	Gender       string             `json:"gender" bson:"gender"`
	ProfilePhoto string             `json:"profile_photo" bson:"profile_photo"`
}

func (m MongoDB) UserModel() *mongo.Collection {
	//return m.Database.Collection("staff")
	return m.Database.Collection("users")
}

func (m MongoDB) Login(u string, p string) (User, error) {
	var user User

	//opts := options.FindOne().SetSort(bson.M{"username": 1})
	//e := m.UserModel().FindOne(context.TODO(), bson.M{"username": u}, opts).Decode(&user)
	// if e == nil {
	// 	e := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p))
	// 	if e == nil {
	// 		return user, nil
	// 	}
	// }

	e := m.UserModel().FindOne(context.Background(), bson.M{"email": u}).Decode(&user)
	if e != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if e == mongo.ErrNoDocuments {
			return user, errors.New("email or password not matching")
		}
		log.Println(e)
	}
	if e == nil && user.Password == p {
		return user, nil
	}
	return user, errors.New("email or password not matching")
}

func (m MongoDB) AuthUser(auth *sessions.Session) (User, error) {
	hexString := fmt.Sprintf("%v", auth.Values["user_id"])
	objID, err := primitive.ObjectIDFromHex(hexString)
	if err != nil {
		log.Println(err)
	}

	var user User
	e := m.UserModel().FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	if e == nil {
		return user, nil
	}
	return user, errors.New("user not found")
}
