package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func ApiHandler() {

	// Automatic Marshal and UnMarshal
	route := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Middleware
	// Cross-Origin Resource Sharing
	route.Use(cors.New())

	// API metrics endpoint
	route.Get("/metrics", monitor.New(monitor.Config{Title: "API Gateway Monitoring Dashboard"}))

	// Endpoints
	route.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is Healthy✅")
	})

	route.Listen(":8085")
}
