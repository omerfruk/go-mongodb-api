package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-mongodb/models"
	"go-mongodb/service"
)

func GetHobbies(c *fiber.Ctx) error {
	hobbies, err := service.GetHobbies()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(hobbies)
}

func GetHobbyByID(c *fiber.Ctx) error {
	id := c.Params("id")
	hobby, err := service.GetHobbyById(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(hobby)
}

func GetUsersHobbies(c *fiber.Ctx) error {
	hobby := c.Params("hobby")
	hobbies, err := service.FindHobbiesUsers(hobby)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(hobbies)
}

func CreateHobby(c *fiber.Ctx) error {
	var hobby models.Hobby
	err := c.BodyParser(&hobby)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	err = service.CreateHobby(hobby)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).SendString("Hobby created")
}

func UpdateHobby(c *fiber.Ctx) error {
	id := c.Params("id")
	var hobby models.Hobby
	err := c.BodyParser(&hobby)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	err = service.UpdateHobby(id, hobby)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).SendString("Hobby updated")
}

func DeleteHobby(c *fiber.Ctx) error {
	id := c.Params("id")
	err := service.DeleteHobby(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).SendString("Hobby deleted")
}
