package main

import (
	"Airplane-Reservation/config"
	"Airplane-Reservation/controllers"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.BodyLimit("1M"))
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())

	controllers.Setup(e)

	e.Start(":" + strconv.Itoa(config.Port))
}
