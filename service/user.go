package service

import (
	"context"
	"fmt"
	"go-mongodb/database"
	"go-mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(m models.User) {
	result, err := database.UserCollection.InsertOne(context.TODO(), m)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.InsertedID)
}

func GetUserById(id string) models.User {
	var user models.User
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	err = database.UserCollection.FindOne(database.Ctx, bson.D{{"_id", objectId}}).
		Decode(&user)
	if err != nil {
		panic(err)
	}
	return user
}

func GetUsers() ([]models.User, error) {
	var user models.User
	var users []models.User
	cur, err := database.UserCollection.Find(database.Ctx, bson.D{})
	if err != nil {
		defer cur.Close(database.Ctx)
		return nil, err
	}
	for cur.Next(database.Ctx) {
		err = cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
