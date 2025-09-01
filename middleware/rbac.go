package middleware

import (
	"myfiberapi/database"
	"myfiberapi/models"

	"github.com/gofiber/fiber/v2"
)

func RequireRole(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userID").(uint)
		
		var user models.User
		if err := database.DB.First(&user, userID).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status": "error",
				"message": "User not found",
			})
		}

		// Check if user's role is in the allowed roles
		isAllowed := false
		for _, role := range roles {
			if user.Role == role {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"status": "error",
				"message": "Access denied: insufficient privileges",
			})
		}

		return c.Next()
	}
}
