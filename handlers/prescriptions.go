package handlers

import (
	"github.com/Rolas444/goapi.git/database"
	"github.com/Rolas444/goapi.git/models"
	"github.com/gofiber/fiber/v2"
)

func CreatePrescription(c *fiber.Ctx) error {
	prescription := new(models.Prescription)
	if err := c.BodyParser(prescription); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DB.Create(&prescription)
	return c.JSON(prescription)
}

func GetPrescription(c *fiber.Ctx) error {
	id := c.Params("id")
	var prescription models.Prescription
	database.DB.Where("id_person = ?", id).First(&prescription)
	return c.JSON(prescription)
}
