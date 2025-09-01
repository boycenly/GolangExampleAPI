package routes

import (
	"myfiberapi/database"
	"myfiberapi/middleware"
	"myfiberapi/models"
	"myfiberapi/utils"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	api := app.Group("/api/users", middleware.JWTProtected())

	// Get all users (admin only)
	api.Get("/", middleware.RequireRole("admin"), func(c *fiber.Ctx) error {
		var users []models.User
		if err := database.DB.Find(&users).Error; err != nil {
			return utils.Error(c, fiber.StatusInternalServerError, "Failed to get users")
		}

		// Map ke response tanpa password
		var responses []UserResponse
		for _, u := range users {
			responses = append(responses, UserResponse{
				ID:    u.ID,
				Name:  u.Name,
				Email: u.Email,
				Role:  u.Role,
			 
			})
		}

		return utils.Success(c, "Users retrieved successfully", responses)
	})

	// Get user by ID (admin can get any user, normal user can only get their own)
	api.Get("/:id", func(c *fiber.Ctx) error {
		userID := c.Locals("userID").(uint)
		requestedID, err := c.ParamsInt("id")
		if err != nil {
			return utils.Error(c, fiber.StatusBadRequest, "Invalid user ID")
		}

		// Check if user is trying to access other's data
		if uint(requestedID) != userID {
			// If not admin, deny access
			var currentUser models.User
			if err := database.DB.First(&currentUser, userID).Error; err != nil {
				return utils.Error(c, fiber.StatusUnauthorized, "User not found")
			}
			if currentUser.Role != "admin" {
				return utils.Error(c, fiber.StatusForbidden, "Access denied")
			}
		}

		var user models.User
		if err := database.DB.First(&user, requestedID).Error; err != nil {
			return utils.Error(c, fiber.StatusNotFound, "User not found")
		}

		return utils.Success(c, "User retrieved successfully", UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		})
	})

	// Update user (admin can update any user, normal user can only update themselves)
	api.Put("/:id", func(c *fiber.Ctx) error {
		userID := c.Locals("userID").(uint)
		requestedID, err := c.ParamsInt("id")
		if err != nil {
			return utils.Error(c, fiber.StatusBadRequest, "Invalid user ID")
		}

		// Check if user is trying to update other's data
		if uint(requestedID) != userID {
			// If not admin, deny access
			var currentUser models.User
			if err := database.DB.First(&currentUser, userID).Error; err != nil {
				return utils.Error(c, fiber.StatusUnauthorized, "User not found")
			}
			if currentUser.Role != "admin" {
				return utils.Error(c, fiber.StatusForbidden, "Access denied")
			}
		}

		var updateReq struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		if err := c.BodyParser(&updateReq); err != nil {
			return utils.Error(c, fiber.StatusBadRequest, "Invalid request body")
		}

		var user models.User
		if err := database.DB.First(&user, requestedID).Error; err != nil {
			return utils.Error(c, fiber.StatusNotFound, "User not found")
		}

		// Update fields
		if updateReq.Name != "" {
			user.Name = updateReq.Name
		}
		if updateReq.Email != "" {
			user.Email = updateReq.Email
		}

		if err := database.DB.Save(&user).Error; err != nil {
			return utils.Error(c, fiber.StatusInternalServerError, "Failed to update user")
		}

		return utils.Success(c, "User updated successfully", UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		})
	})

	// Delete user (admin only)
	api.Delete("/:id", middleware.RequireRole("admin"), func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return utils.Error(c, fiber.StatusBadRequest, "Invalid user ID")
		}

		var user models.User
		if err := database.DB.First(&user, id).Error; err != nil {
			return utils.Error(c, fiber.StatusNotFound, "User not found")
		}

		if err := database.DB.Delete(&user).Error; err != nil {
			return utils.Error(c, fiber.StatusInternalServerError, "Failed to delete user")
		}

		return utils.Success(c, "User deleted successfully", nil)
	})
}
