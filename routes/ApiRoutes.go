package routes

import (
	"github.com/barankibar/Rick-Morty-Rest-API/routes/controllers"
	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app *fiber.App) {
	app.Post("/api/v1/create/character", controllers.CreateChar)
	app.Post("/api/v1/characters", controllers.CreateCharacters)

	app.Get("/api/character/:id", controllers.GetACharacter)
}
