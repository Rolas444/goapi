package handlers

import (
	"github.com/Rolas444/goapi.git/database"
	"github.com/Rolas444/goapi.git/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllPersons(c *fiber.Ctx) error {
	var persons []models.Person
	database.DB.Find(&persons)
	return c.JSON(persons)
}

func CreatePerson(c *fiber.Ctx) error {
	person := new(models.Person)
	if err := c.BodyParser(person); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DB.Create(&person)
	return c.JSON(person)
}

func GetPerson(c *fiber.Ctx) error {
	id := c.Params("id")
	var person models.Person
	database.DB.Find(&person, id)
	return c.JSON(person)
}

func UpdatePerson(c *fiber.Ctx) error {

	id := c.Params("id")
	person := new(models.Person)
	if err := c.BodyParser(person); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	var oldPerson models.Person
	database.DB.Find(&oldPerson, id)
	if oldPerson.ID == 0 {
		return c.Status(404).SendString("No person found with ID")
	}
	database.DB.Model(&oldPerson).Updates(person)
	return c.JSON(oldPerson)
}
