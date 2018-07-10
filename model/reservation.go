package model

import "Airplane-Reservation/db"

type ReservationBase struct {
	Airplane Airplane
	Booker   User
}

type Reservation struct {
	ReservationBase
}

type WaitingReservation struct {
	ReservationBase
	Timestamp int
}

var (
	ReservationCol        = db.DB().C("reservations")
	ReservationWaitingCol = db.DB().C("reservation_waitings")
)
