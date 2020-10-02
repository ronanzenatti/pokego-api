package api

import (
	"antares-api/api/handler"
	"antares-api/config"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Server - Function that starts the echo server.
func Server() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/server/status", func(context echo.Context) error {
		return handler.ServerStatus.GetStatus(context)
	})

	e.GET("/pokemon/:pokemon", func(context echo.Context) error {
		return handler.Pokemon.GetByName(context)
	})

	e.Logger.Fatal(e.Start(":" + config.APIPort))
}
