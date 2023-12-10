package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waldeedle/reporting/internal/private_api"
	"github.com/waldeedle/reporting/internal/public_api"
)

func main() {
	app := fiber.New()

	err := public_api.AddRoutes(app)
	if err != nil {
		panic(err)
	}

	err = private_api.AddRoutes(app)
	if err != nil {
		panic(err)
	}

	app.Listen(":3000")
}
