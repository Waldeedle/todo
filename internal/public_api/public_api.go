package public_api

import "github.com/gofiber/fiber/v2"

func AddRoutes(app *fiber.App) error {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	return nil
}
