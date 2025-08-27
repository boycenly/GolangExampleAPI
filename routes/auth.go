package routes

import (
	"strings"

	"myfiberapi/database"
	"myfiberapi/models"
	"myfiberapi/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type registerReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SetupAuthRoutes(app *fiber.App) {
	g := app.Group("/auth")

	g.Post("/register", func(c *fiber.Ctx) error {
		var req registerReq
		if err := c.BodyParser(&req); err != nil {
			return utils.Error(c, fiber.StatusBadRequest, "Invalid request body")
		}
		req.Email = strings.TrimSpace(strings.ToLower(req.Email))
		if req.Name == "" || req.Email == "" || len(req.Password) < 6 {
			return utils.Error(c, fiber.StatusBadRequest, "Name, email, and password(>=6) required")
		}
		// cek exist
		var count int64
		database.DB.Model(&models.User{}).Where("email = ?", req.Email).Count(&count)
		if count > 0 {
			return utils.Error(c, fiber.StatusConflict, "Email already registered")
		}
		// hash pass
		hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		u := models.User{Name: req.Name, Email: req.Email, Password: string(hash)}
		if err := database.DB.Create(&u).Error; err != nil {
			return utils.Error(c, fiber.StatusInternalServerError, "Failed to create user")
		}
		return utils.Success(c, "Registered successfully", fiber.Map{
			"id": u.ID, "name": u.Name, "email": u.Email,
		})
	})

	g.Post("/login", func(c *fiber.Ctx) error {
		var req loginReq
		if err := c.BodyParser(&req); err != nil {
			return utils.Error(c, fiber.StatusBadRequest, "Invalid request body")
		}
		req.Email = strings.TrimSpace(strings.ToLower(req.Email))
		var u models.User
		if err := database.DB.Where("email = ?", req.Email).First(&u).Error; err != nil {
			return utils.Error(c, fiber.StatusUnauthorized, "Invalid email or password")
		}
		if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
			return utils.Error(c, fiber.StatusUnauthorized, "Invalid email or password")
		}
		token, err := utils.GenerateToken(u.ID)
		if err != nil {
			return utils.Error(c, fiber.StatusInternalServerError, "Failed to generate token")
		}
		return utils.Success(c, "Login success", fiber.Map{
			"token": token,
			"user":  fiber.Map{"id": u.ID, "name": u.Name, "email": u.Email},
		})
	})
}
