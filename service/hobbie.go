package service

import (
	"context"
	"fmt"
	"go-mongodb/database"
	"go-mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateHobbie(m models.Hobbie) {
	result, err := database.HobbieCollection.InsertOne(context.TODO(), m)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.InsertedID)
}

func GetHobbieById(id string) models.Hobbie {
	var hobbie models.Hobbie
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	err = database.HobbieCollection.FindOne(database.Ctx, bson.D{{"_id", objectId}}).
		Decode(&hobbie)
	if err != nil {
		panic(err)
	}
	return hobbie
}

func GetHobbies() ([]models.Hobbie, error) {
	var hobbie models.Hobbie
	var hobbies []models.Hobbie
	cur, err := database.HobbieCollection.Find(database.Ctx, bson.D{})
	if err != nil {
		defer cur.Close(database.Ctx)
		return nil, err
	}
	for cur.Next(database.Ctx) {
		err = cur.Decode(&hobbie)
		if err != nil {
			return nil, err
		}
		hobbies = append(hobbies, hobbie)
	}
	return hobbies, nil
}

func UpdateHobbie(id string, hobbi models.Hobbie) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$set", bson.D{{"name", hobbi.Name}}}}
	_, err = database.HobbieCollection.UpdateOne(
		database.Ctx,
		filter,
		update,
	)
	return err
}

func DeleteHobbie(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = database.HobbieCollection.DeleteOne(database.Ctx, bson.D{{"_id", objectId}})
	if err != nil {
		return err
	}
	return nil
}
