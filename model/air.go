package model

import "Airplane-Reservation/db"

type Airline struct {
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
