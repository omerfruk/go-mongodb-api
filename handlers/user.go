package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-mongodb/models"
	"go-mongodb/service"
)

func GetUsers(c *fiber.Ctx) error {
	users, err := service.GetUsers()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(users)
}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := service.GetUserById(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	err = service.CreateUser(user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).SendString("User created")
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	err = service.UpdateUser(id, user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).SendString("User updated")
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := service.DeleteUser(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).SendString("User deleted")
}
