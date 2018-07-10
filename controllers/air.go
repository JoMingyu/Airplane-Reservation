package controllers

import (
	"Airplane-Reservation/model"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2-unstable/bson"
)

func GetAirlines(c echo.Context) error {
	airlines := []model.Airline{}

	model.AirlineCol.Find(bson.M{}).All(&airlines)

	return c.JSON(http.StatusOK, airlines)
}
