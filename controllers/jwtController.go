package controllers

import (
	"net/http"
	"time"

	"github.com/barankibar/Rick-Morty-Rest-API/routes/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenStr, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// func VerifyJWT(token string) (bool, error) {
// 	claims := &Claims{}

// 	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	if err != nil {
// 		return false, err
// 	}
// 	if !tkn.Valid {
// 		return false, err
// 	}

// 	return true, nil

// }

func HandleGenerateToken(c *fiber.Ctx) error {
	token, err := CreateJWT()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.JWTCreatedResponse{Status: 400, Message: "Token Couldn't Created", Data: &fiber.Map{"data": err.Error()}})
	}
	return c.Status(http.StatusCreated).JSON(responses.JWTCreatedResponse{Status: http.StatusCreated, Message: "Token-Created", Algorithm: "HS256", Data: &fiber.Map{"data": token}})
}
