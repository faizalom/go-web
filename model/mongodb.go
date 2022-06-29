package model

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	*mongo.Database
}

func ConnDB() *mongo.Database {
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://faizalom:Faiare123@cluster0.urcqt.mongodb.net/next_demo?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//return client.Database("next_demo")
	return client.Database("next_demo")
}
