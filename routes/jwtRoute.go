package routes

import (
	"github.com/barankibar/Rick-Morty-Rest-API/routes/controllers"
	"github.com/gofiber/fiber/v2"
)

func JwtRoutes(app *fiber.App) {
	app.Get("/generate-token", controllers.HandleGenerateToken)
}
