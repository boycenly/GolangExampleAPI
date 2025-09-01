package routes

import (
	"myfiberapi/middleware"
	"myfiberapi/models"
	"myfiberapi/utils"

	"github.com/gofiber/fiber/v2"
)

func SetupExampleRoutes(app *fiber.App) {
	api := app.Group("/api", middleware.JWTProtected())

	// Route yang bisa diakses semua user yang sudah login
	api.Get("/hello", func(c *fiber.Ctx) error {
		msg := models.Message{Text: "Hello World (protected)"}
		return c.JSON(msg)
	})

	// Route yang hanya bisa diakses admin
	api.Get("/admin-only", middleware.RequireRole("admin"), func(c *fiber.Ctx) error {
		return utils.Success(c, "Success", fiber.Map{
			"message": "This is admin only route",
		})
	})

	// Route yang bisa diakses admin atau manager
	api.Get("/manager-area", middleware.RequireRole("admin", "manager"), func(c *fiber.Ctx) error {
		return utils.Success(c, "Success", fiber.Map{
			"message": "This is manager or admin area",
		})
	})

	api.Get("/user-only", middleware.RequireRole("user"), func(c *fiber.Ctx) error {
		return utils.Success(c, "Success", fiber.Map{
			"message": "This is user only route",
		})
	})
}
