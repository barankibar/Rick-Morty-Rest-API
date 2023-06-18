package routes

import (
	"github.com/barankibar/Rick-Morty-Rest-API/routes/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Post("/login", controllers.UserLogin)
	app.Post("/register", controllers.CreateUser)
}
