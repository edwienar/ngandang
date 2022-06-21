package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Framework echo is working!")
	})

	return e
}
