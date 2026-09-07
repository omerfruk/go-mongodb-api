package service

import (
	"github.com/omerfruk/go-mongodb-api/database"
	"github.com/omerfruk/go-mongodb-api/models"
	"github.com/omerfruk/go-mongodb-api/viewmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateHobby(m models.Hobby) error {
	_, err := database.HobbyCollection.InsertOne(database.Ctx, m)
	return err
}

func GetHobbyById(id string) (models.Hobby, error) {
	var Hobby models.Hobby
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Hobby, err
	}
	err = database.HobbyCollection.FindOne(database.Ctx, bson.D{{"_id", objectId}}).
		Decode(&Hobby)
	return Hobby, err
}

func GetHobbies() ([]models.Hobby, error) {
	var Hobby models.Hobby
	var hobbies []models.Hobby
	cur, err := database.HobbyCollection.Find(database.Ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(database.Ctx)
	for cur.Next(database.Ctx) {
		err = cur.Decode(&Hobby)
		if err != nil {
			return nil, err
		}
		hobbies = append(hobbies, Hobby)
	}
	return hobbies, cur.Err()
}

func UpdateHobby(id string, hobbi models.Hobby) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$set", bson.D{{"name", hobbi.Name}}}}
	_, err = database.HobbyCollection.UpdateOne(
		database.Ctx,
		filter,
		update,
	)
	return err
}

func DeleteHobby(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = database.HobbyCollection.DeleteOne(database.Ctx, bson.D{{"_id", objectId}})
	if err != nil {
		return err
	}
	return nil
}

func FindHobbiesUsers(HobbyName string) ([]models.User, error) {
	matchStage := bson.D{{"$match", bson.D{{"name", HobbyName}}}}

	lookupStage := bson.D{{"$lookup",
		bson.D{{"from", "users"},
			{"localField", "name"},
			{"foreignField", "hobby"},
			{"as", "users"}}}}

	showLoadedCursor, err := database.HobbyCollection.Aggregate(database.Ctx,
		mongo.Pipeline{matchStage, lookupStage})
	if err != nil {
		return nil, err
	}
	defer showLoadedCursor.Close(database.Ctx)

	var returnModel []viewmodel.HobbiesUsers
	if err := showLoadedCursor.All(database.Ctx, &returnModel); err != nil {
		return nil, err
	}
	if len(returnModel) == 0 {
		return []models.User{}, nil
	}

	return returnModel[0].Users, nil
}
