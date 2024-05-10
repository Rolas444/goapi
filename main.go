package main

import (
	"log"

	"github.com/Rolas444/goapi.git/database"
	"github.com/Rolas444/goapi.git/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDB()

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		// Define the maximum cookie size
		maxCookieSize := 4096

		// Get the Cookie header
		cookieHeader := c.Get("Cookie")

		// Check the size of the Cookie header
		if len(cookieHeader) > maxCookieSize {
			// If the Cookie header is too large, return an error
			return c.Status(431).SendString("Request Header Fields Too Large")
		}

		// If the Cookie header is not too large, continue with the next handler
		return c.Next()
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: false,
		AllowHeaders:     "Origin, Content-Type, Accept, Accept-Encoding",
	}))

	// app.Use(compress.New())
	app.Use(func(c *fiber.Ctx) error {
		// Clear all cookies
		c.Request().Header.Del("Cookie")

		// Continue with the next handler
		return c.Next()
	})

	// app.Use(func(c *fiber.Ctx) error {

	// 	maxHeaderSize := 4096
	// 	headerSize := 0
	// 	c.Request().Header.VisitAll(func(key, value []byte) {
	// 		headerSize += len(key) + len(value)
	// 	})
	// 	if headerSize > maxHeaderSize {
	// 		return c.Status(431).SendString("Request Header Fields Too Large 2")
	// 	}

	// 	return c.Next()
	// })

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":8000"))

}
