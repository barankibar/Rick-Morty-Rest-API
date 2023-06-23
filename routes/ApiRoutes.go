package routes

import (
	"github.com/barankibar/Rick-Morty-Rest-API/routes/controllers"
	"github.com/barankibar/Rick-Morty-Rest-API/routes/middlewares"
	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app *fiber.App) {
	api := app.Group("/api/v1", middlewares.JWTProtected())

	api.Post("/create/character", controllers.CreateChar)
	api.Post("/characters", controllers.CreateCharacters)

	api.Get("/character/:id", controllers.GetACharacter)
}
