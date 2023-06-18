package responses

import "github.com/gofiber/fiber/v2"

type JWTCreatedResponse struct {
	Status    int        `json:"status"`
	Message   string     `json:"message"`
	Algorithm string     `json:"alg"`
	Type      string     `json:"typ"`
	Data      *fiber.Map `json:"data"`
}
