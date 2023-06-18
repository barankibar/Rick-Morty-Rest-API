package routes

import (
	"github.com/barankibar/Rick-Morty-Rest-API/routes/controllers"
	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app *fiber.App) {
	app.Post("/api/create/character", controllers.CreateChar)
	app.Post("/api/characters", controllers.CreateCharacters)

	app.Get("/api/user/:userId", controllers.GetACharacter)
}
