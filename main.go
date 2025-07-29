package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!!!")
	})

	app.Get("/json", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Username": "Aditya Ibrar Abdillah ",
			"Role":     "Mobile Developer",
		})
	})

	app.Listen(":3000")
}
