package private_api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waldeedle/reporting/internal/middleware"
)

func AddRoutes(app *fiber.App) error {
	middleware.AddPPROFToApp(app)

	app.Get("/tools", func(c *fiber.Ctx) error {
		return c.SendString("Shhhhh Secret Tools 🤫!")
	})

	return nil
}
