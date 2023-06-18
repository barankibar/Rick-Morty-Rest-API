package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/barankibar/Rick-Morty-Rest-API/routes/configs"
	"github.com/barankibar/Rick-Morty-Rest-API/routes/models"
	"github.com/barankibar/Rick-Morty-Rest-API/routes/responses"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var charCollection *mongo.Collection = configs.GetCollection(configs.DB, "characters")
var validate = validator.New()

func GetACharacter(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	charID := c.Params("id")

	var character models.Character

	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(charID)

	err := charCollection.FindOne(ctx, bson.M{"id": objID}).Decode(&character)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.CharResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.CharResponse{Status: http.StatusOK, Message: "OK", Data: &fiber.Map{"data": character}})

}

func CreateChar(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var character models.Character
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&character); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.CharResponse{Status: http.StatusBadRequest, Message: "Error", Data: &fiber.Map{"data": err.Error()}})
	}
	//use the validator library to validate required fields
	if validationErr := validate.Struct(&character); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.CharResponse{Status: http.StatusBadRequest, Message: "Error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newChar := models.Character{
		ID:       primitive.NewObjectID(),
		Name:     character.Name,
		Status:   character.Status,
		Species:  character.Status,
		Gender:   character.Gender,
		Origin:   character.Origin,
		Location: character.Location,
		Image:    character.Image,
		Episode:  character.Episode,
		URL:      character.URL,
		Created:  character.Created,
	}

	result, err := charCollection.InsertOne(ctx, newChar)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.CharResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.CharResponse{Status: http.StatusCreated, Message: "Success", Data: &fiber.Map{"data": result}})
}

func CreateCharacters(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var payload struct {
		Results []models.Character `json:"results"`
	}
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.CharResponse{Status: http.StatusBadRequest, Message: "Error", Data: &fiber.Map{"data": err.Error()}})
	}
	var documents []interface{}

	for _, character := range payload.Results {
		//use the validator library to validate required fields
		if validationErr := validate.Struct(&character); validationErr != nil {
			return c.Status(http.StatusBadRequest).JSON(responses.CharResponse{Status: http.StatusBadRequest, Message: "Error", Data: &fiber.Map{"data": validationErr.Error()}})
		}

		newChar := models.Character{
			ID:       primitive.NewObjectID(),
			Name:     character.Name,
			Status:   character.Status,
			Species:  character.Status,
			Gender:   character.Gender,
			Origin:   character.Origin,
			Location: character.Location,
			Image:    character.Image,
			Episode:  character.Episode,
			URL:      character.URL,
			Created:  character.Created,
		}
		documents = append(documents, newChar)
	}

	result, err := charCollection.InsertMany(ctx, documents)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.CharResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &fiber.Map{"data": err.Error()}})
	}
	return c.Status(http.StatusCreated).JSON(responses.CharResponse{Status: http.StatusCreated, Message: "Success", Data: &fiber.Map{"data": result}})
}
