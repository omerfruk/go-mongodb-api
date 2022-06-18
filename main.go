package main

import (
	"go-mongodb/database"
	"go-mongodb/service"
)

func main() {
	database.Setup()
	service.GetUsers()
}
