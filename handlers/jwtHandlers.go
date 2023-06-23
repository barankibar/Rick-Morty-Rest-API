package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/barankibar/Rick-Morty-Rest-API/routes/configs"
	"github.com/barankibar/Rick-Morty-Rest-API/routes/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	Expires int64
}

type JWTClaim struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func HandleGenerateToken(user models.User) (string, error) {

	claims := JWTClaim{
		user.UserName,
		user.Password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 4).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate Token
	ss, err := token.SignedString([]byte(configs.EnvJWTSecret()))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println("User is : ", claims["UserName"])
	if ok && token.Valid {
		expires := int64(claims["exp"].(float64))

		return &TokenMetadata{
			Expires: expires,
		}, nil
	}
	return nil, err
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	// Normally Auth Http Header
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}
	return ""
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(configs.EnvJWTSecret()), nil
}
