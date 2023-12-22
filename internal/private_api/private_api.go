package private_api

import (
	"net/http"

	"github.com/labstack/echo-contrib/pprof"
	"github.com/labstack/echo/v4"
)

func AddRoutes(e *echo.Echo) error {
	tools := e.Group("/tools")

	//Add pprof routes
	pprof.Register(e, "/tools/pprof")

	//Add custom routes
	tools.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Shhhhh Secret Tools 🤫!")
	})

	return nil
}
