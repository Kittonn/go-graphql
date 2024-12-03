package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *app) MapHandlers() error {

	a.echo.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "OK",
		})
	})

	return nil
}
