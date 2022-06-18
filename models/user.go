package models

type User struct {
	Name    string `bson:"name"`
	Surname string `bson:"surname"`
	Hobbie  string `bson:"hobbie"`
	Age     int    `bson:"age"`
}
