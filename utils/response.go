package utils

import "github.com/gofiber/fiber/v2"

// Response sukses
func Success(c *fiber.Ctx, message string, data interface{}) error {
    return c.JSON(fiber.Map{
        "status":  "success",
        "message": message,
        "data":    data,
    })
}

// Response error
func Error(c *fiber.Ctx, code int, message string) error {
    return c.Status(code).JSON(fiber.Map{
        "status":  "error",
        "message": message,
    })
}
