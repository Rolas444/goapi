package router

import (
	"github.com/Rolas444/goapi.git/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.Example)
	app.Get("/apigo", handlers.Example)
	app.Get("/persons", handlers.GetAllPersons)
	app.Post("/persons", handlers.CreatePerson)
	app.Put("/persons/:id", handlers.UpdatePerson)

	app.Get("/medications", handlers.GetAllMedications)
	app.Post("/medications", handlers.CreateMedication)
	app.Put("/medications/:id", handlers.UpdateMedication)

	app.Get("/prescriptions/:id", handlers.GetPrescription)
	app.Post("/prescriptions", handlers.CreatePrescription)

	app.Get("/dosages/:id", handlers.GetDosages)
	app.Post("/dosages", handlers.CreateDosage)
	app.Put("/dosages/:id", handlers.UpdateDosage)
	app.Delete("/dosages/:id", handlers.DeleteDosage)
}
