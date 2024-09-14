package main

import (
	database "fiber/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	database.ConnectDb()
	app.Listen(":3000")
}
