package main

import (
	"github.com/barankibar/Rick-Morty-Rest-API/routes/configs"
	"github.com/barankibar/Rick-Morty-Rest-API/routes/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})

	routes.CharRoutes(app)

	app.Listen(":6000")
}
