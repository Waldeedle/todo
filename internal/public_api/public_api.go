package public_api

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/waldeedle/todo/internal/accounts"
	"github.com/waldeedle/todo/internal/todos"
	"github.com/waldeedle/todo/internal/ui"
)

type API struct {
	accounts accounts.Service
	todos    todos.Service
}

func New(accounts accounts.Service, todos todos.Service) API {
	return API{
		accounts: accounts,
		todos:    todos,
	}
}

// todo: add account routes
func (api *API) AddRoutes(e *echo.Echo) error {
	e.Static("/static", "internal/assets")

	e.GET("/", func(c echo.Context) error {
		return HTML(c, ui.Index())
	})

	e.GET("/test2", func(c echo.Context) error {
		return HTML(c, ui.JsTest())
	})

	return nil
}

func HTML(c echo.Context, comp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return comp.Render(c.Request().Context(), c.Response().Writer)
}
