package main

import (
	"github.com/gofiber/fiber/v2"
	"go-mongodb/database"
	"go-mongodb/router"
)

func main() {
	database.Setup()

	app := fiber.New()
	router.Setup(app)
	app.Listen(":3000")
}
