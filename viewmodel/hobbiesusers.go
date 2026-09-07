package viewmodel

import "github.com/omerfruk/go-mongodb-api/models"

type HobbiesUsers struct {
	UserHobbies string        `bson:"name" json:"hobby"`
	Users       []models.User `bson:"users" json:"users"`
}
