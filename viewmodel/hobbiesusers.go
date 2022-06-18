package viewmodel

import "go-mongodb/models"

type HobbiesUsers struct {
	UserHobbies string `bson:"userhobbies"`
	Users       []models.User
}
