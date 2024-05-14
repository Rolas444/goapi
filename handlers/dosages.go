package handlers

import (
	"github.com/Rolas444/goapi.git/database"
	"github.com/Rolas444/goapi.git/models"
	"github.com/gofiber/fiber/v2"
)

func GetDosages(c *fiber.Ctx) error {
	id := c.Params("id")
	var dosages []models.Dosage
	database.DB.Where("prescription_id = ?", id).Find(&dosages)
	return c.JSON(dosages)
}

func CreateDosage(c *fiber.Ctx) error {
	dosage := new(models.Dosage)
	if err := c.BodyParser(dosage); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DB.Create(&dosage)
	return c.JSON(dosage)
}

func UpdateDosage(c *fiber.Ctx) error {
	id := c.Params("id")
	dosage := new(models.Dosage)
	if err := c.BodyParser(dosage); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	var oldDosage models.Dosage
	database.DB.Find(&oldDosage, id)
	if oldDosage.ID == 0 {
		return c.Status(404).SendString("No dosage found with ID")
	}
	database.DB.Model(&oldDosage).Updates(dosage)
	return c.JSON(oldDosage)
}

func DeleteDosage(c *fiber.Ctx) error {
	id := c.Params("id")
	var dosage models.Dosage
	database.DB.Delete(&dosage, id)
	return c.SendString("Dosage successfully deleted")
}
