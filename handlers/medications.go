package handlers

import (
	"github.com/Rolas444/goapi.git/database"
	"github.com/Rolas444/goapi.git/models"
	"github.com/gofiber/fiber/v2"
)

func Example(c *fiber.Ctx) error {
	return c.SendString("hello world")
}

func GetAllMedications(c *fiber.Ctx) error {
	var medications []models.Medication
	database.DB.Find(&medications)
	return c.JSON(medications)
}

func CreateMedication(c *fiber.Ctx) error {
	medication := new(models.Medication)
	if err := c.BodyParser(medication); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DB.Create(&medication)
	return c.JSON(medication)
}

func GetMedication(c *fiber.Ctx) error {
	id := c.Params("id")
	var medication models.Medication
	database.DB.Find(&medication, id)
	return c.JSON(medication)
}

func UpdateMedication(c *fiber.Ctx) error {
	id := c.Params("id")
	medication := new(models.Medication)
	if err := c.BodyParser(medication); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	var oldMedication models.Medication
	database.DB.Find(&oldMedication, id)
	if oldMedication.ID == 0 {
		return c.Status(404).SendString("No medication found with ID")
	}
	database.DB.Model(&oldMedication).Updates(medication)
	return c.JSON(oldMedication)
}
