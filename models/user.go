package models

type User struct {
	Name    string `bson:"name"`
	Surname string `bson:"surname"`
	Hobby   string `bson:"hobbie"`
	Age     int    `bson:"age"`
}
