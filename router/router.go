package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/omerfruk/go-mongodb-api/handlers"
)

func Setup(app fiber.Router) {
	app.Use(logger.New())

	app.Get("/users", handlers.GetUsers)
	app.Get("/users/:id", handlers.GetUserByID)
	app.Post("/users", handlers.CreateUser)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.DeleteUser)

	app.Get("/hobbies", handlers.GetHobbies)
	app.Get("/hobbies/:id", handlers.GetHobbyByID)
	app.Post("/hobbies", handlers.CreateHobby)
	app.Put("/hobbies/:id", handlers.UpdateHobby)
	app.Delete("/hobbies/:id", handlers.DeleteHobby)
	app.Get("/users/hobbies/:hobby", handlers.GetUsersHobbies)
}
