package viewmodel

import "go-mongodb/models"

type HobbiesUsers struct {
	UserHobbies string `bson:"name"`
	Users       []models.User
}
