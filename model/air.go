package model

import (
	"Airplane-Reservation/db"

	"gopkg.in/mgo.v2-unstable/bson"
)

type Airline struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string
}

type Airplane struct {
	DepartmentTimeHour  int
	SeatCount           int
	ReservableSeatCount int
	Airline             Airline
}

var (
	AirlineCol  = db.DB().C("airlines")
	AirplaneCol = db.DB().C("airplanes")
)
