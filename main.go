package main

import (
	"api-event-ticket/configs"
	"api-event-ticket/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})
	configs.ConnectDB()
	routes.UserRoute(app)

	app.Listen(":4000")
}
