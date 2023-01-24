package handlers

import "github.com/gofiber/fiber/v2"

func ApiHandler() {

	route := fiber.New()

	route.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is Healthy✅")
	})

	route.Listen(":8085")
}
