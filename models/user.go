package models

type User struct {
	Name    string `bson:"name"`
	Surname string `bson:"surname"`
	Age     int    `bson:"age"`
}
