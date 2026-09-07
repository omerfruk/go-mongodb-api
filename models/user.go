package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name    string             `bson:"name" json:"name"`
	Surname string             `bson:"surname" json:"surname"`
	Hobby   string             `bson:"hobby" json:"hobby"`
	Age     int                `bson:"age" json:"age"`
}
