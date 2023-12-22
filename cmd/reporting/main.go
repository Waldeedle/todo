package main

import (
	"github.com/labstack/echo/v4"
	"github.com/waldeedle/reporting/internal/private_api"
	"github.com/waldeedle/reporting/internal/public_api"
)

func main() {
	e := echo.New()

	err := public_api.AddRoutes(e)
	if err != nil {
		panic(err)
	}

	err = private_api.AddRoutes(e)
	if err != nil {
		panic(err)
	}

	e.Logger.Fatal(e.Start(":3000"))
}
