package router

import (
	"github.com/Rolas444/goapi.git/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.Example)
	app.Get("/apigo", handlers.Example)
	app.Get("/persons", handlers.GetAllPersons)

	app.Get("/medications", handlers.GetAllMedications)
}
