package routes

import (
	"myfiberapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupExampleRoutes(app *fiber.App) {
	app.Get("/hello", func(c *fiber.Ctx) error {
		msg := models.Message{
			Text: "Hello World",
		}
		return c.JSON(msg)
	})
}
