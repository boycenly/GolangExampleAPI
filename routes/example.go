package routes

import (
	"myfiberapi/middleware"
	"myfiberapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupExampleRoutes(app *fiber.App) {
	api := app.Group("/api", middleware.JWTProtected())

	api.Get("/hello", func(c *fiber.Ctx) error {
		msg := models.Message{Text: "Hello World (protected)"}
		return c.JSON(msg)
	})
}
