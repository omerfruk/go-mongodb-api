package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connection URI
const uri = "mongodb://localhost:27017"

var (
	UserCollection  *mongo.Collection
	HobbyCollection *mongo.Collection
	Ctx             = context.TODO()
)

func Setup() {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	// Ping the primary
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	db := client.Database("go-mongodb")

	UserCollection = db.Collection("users")
	HobbyCollection = db.Collection("hobbie")
}
