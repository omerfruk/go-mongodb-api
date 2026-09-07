package service

import (
	"github.com/omerfruk/go-mongodb-api/database"
	"github.com/omerfruk/go-mongodb-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(m models.User) error {
	_, err := database.UserCollection.InsertOne(database.Ctx, m)
	return err
}

func GetUserById(id string) (models.User, error) {
	var user models.User
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	err = database.UserCollection.FindOne(database.Ctx, bson.D{{"_id", objectId}}).
		Decode(&user)
	return user, err
}

func GetUsers() ([]models.User, error) {
	var user models.User
	var users []models.User
	cur, err := database.UserCollection.Find(database.Ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(database.Ctx)
	for cur.Next(database.Ctx) {
		err = cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, cur.Err()
}

func UpdateUser(id string, user models.User) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$set", bson.M{
		"name":    user.Name,
		"surname": user.Surname,
		"hobby":   user.Hobby,
		"age":     user.Age,
	}}}
	_, err = database.UserCollection.UpdateOne(
		database.Ctx,
		filter,
		update,
	)
	return err
}

func DeleteUser(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = database.UserCollection.DeleteOne(database.Ctx, bson.D{{"_id", objectId}})
	if err != nil {
		return err
	}
	return nil
}
