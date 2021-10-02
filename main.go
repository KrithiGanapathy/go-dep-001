package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	var app = fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("It works 👍")
	})
	app.Listen(":8080")
}
