package utils

import "github.com/gofiber/fiber/v2"

// ApiResponse format umum
type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success response
func Success(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(ApiResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// Error response
func Error(c *fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(ApiResponse{
		Status:  "error",
		Message: message,
	})
}
