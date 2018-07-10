package controllers

import "github.com/labstack/echo"

func Setup(e *echo.Echo) {
	e.GET("/check/id/:id", IDDuplicationCheck)
}
