package routes

import (
	"first_api_golang/controllers"
	"first_api_golang/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB)  {
	authController := &controllers.AuthController{DB: db}

	// Grup route untuk API
	api := app.Group("/api")

	// Route autentikasi
	api.Post("/register", authController.Register)
	api.Post("/login", authController.Login)

	// Route yang dilindungi autentikasi
	protected := api.Group("/protected")
	protected.Use(utils.AuthMiddleware)
	protected.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "This is a protected route",
		})
	})
}