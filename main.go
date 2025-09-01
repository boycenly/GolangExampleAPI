package main

import (
	"fmt"
	"myfiberapi/database"
	"myfiberapi/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// 1. Koneksi DB
	database.Connect()

	// 2. Setup Fiber
	app := fiber.New()
	routes.SetupAuthRoutes(app)
	// 3. Setup routes
	routes.SetupExampleRoutes(app)
	// 4. Setup routes user
	routes.SetupUserRoutes(app)

	// 4. Jalankan servers
	fmt.Println("Server running at http://localhost:3000")
	app.Listen(":3000")
}
