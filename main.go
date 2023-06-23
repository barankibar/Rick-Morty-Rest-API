package main

import (
	"github.com/barankibar/Rick-Morty-Rest-API/routes/configs"
	"github.com/barankibar/Rick-Morty-Rest-API/routes/middlewares"
	"github.com/barankibar/Rick-Morty-Rest-API/routes/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	middlewares.JWTProtected()

	app.Static("/", "./static", fiber.Static{
		Compress: true,
		Browse:   true,
		MaxAge:   3600,
		Index:    "index.html",
	})
	app.Static("/login", "./static", fiber.Static{
		Compress: true,
		Browse:   true,
		MaxAge:   3600,
		Index:    "login.html",
	})
	app.Static("/register", "./static", fiber.Static{
		Compress: true,
		Browse:   true,
		MaxAge:   3600,
		Index:    "register.html",
	})

	// Public Routes
	routes.UserRoutes(app)

	// Private Routes
	routes.ApiRoutes(app)

	app.Listen(":3000")
}
