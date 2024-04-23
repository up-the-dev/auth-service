package application

import "github.com/gofiber/fiber/v2"

func Get_instance() *fiber.App {
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "pong",
		})
	})
	return app
}
