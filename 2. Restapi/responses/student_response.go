package responses

import "github.com/gofiber/fiber/v2"

// The structure of the response which will be sent back to the user

type StudentResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}
