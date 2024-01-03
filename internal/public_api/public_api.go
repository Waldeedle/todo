package public_api

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/waldeedle/todo/internal/templates"
)

func HTML(c echo.Context, comp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return comp.Render(c.Request().Context(), c.Response().Writer)
}

func AddRoutes(e *echo.Echo) error {
	e.Static("/static", "internal/assets")

	e.GET("/", func(c echo.Context) error {
		return HTML(c, templates.Page())
	})

	return nil
}
