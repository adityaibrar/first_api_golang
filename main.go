package main

import (
	"first_api_golang/config"
	"first_api_golang/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(
		cors.Config{
			AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		},
	))

	db := config.InitDatabase()

	routes.SetupRoutes(app, db)

	log.Fatal(app.Listen(":3000"))
}
