package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/barankibar/Rick-Morty-Rest-API/routes/configs"
	"github.com/barankibar/Rick-Morty-Rest-API/routes/handlers"
	"github.com/barankibar/Rick-Morty-Rest-API/routes/models"
	"github.com/barankibar/Rick-Morty-Rest-API/routes/responses"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func UserLogin(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user := new(models.User)
	body := new(models.User)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	err := userCollection.FindOne(ctx, bson.D{primitive.E{Key: "username", Value: body.UserName}}).Decode(&user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.CharResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &fiber.Map{"data": err.Error()}})
	}

	if body.Password != user.Password {
		return c.Status(http.StatusUnauthorized).JSON(responses.UserResponse{Message: "Error", Status: http.StatusUnauthorized, Data: &fiber.Map{"data": "Invalid Password"}})
	}

	t, err := handlers.HandleGenerateToken(*user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.JWTCreatedResponse{Message: "Error", Status: http.StatusInternalServerError, Data: &fiber.Map{"Error": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.JWTCreatedResponse{Message: "JWT Created", Status: http.StatusOK, Data: &fiber.Map{"access_token": t}, Algorithm: "HS256", Type: "JWT"})
}

func CreateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User

	defer cancel()

	// Validate request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "Error", Data: &fiber.Map{"data": err.Error()}})
	}
	// Validation Error
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "Error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newUser := models.User{
		UserName: user.UserName,
		Password: user.Password,
	}

	err := userCollection.FindOne(ctx, bson.D{primitive.E{Key: "username", Value: user.UserName}}).Decode(&user)
	if err == nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "Error : Username must be unique", Data: &fiber.Map{"data": "username must be unique"}})
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "Success", Data: &fiber.Map{"data": result}})
}
