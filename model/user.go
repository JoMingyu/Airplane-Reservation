package model

import "Airplane-Reservation/db"

type User struct {
	ID   string `bson:"_id"`
	Pw   string
	Name string
}

type Token struct {
	Identity string `bson:"_id"`
	Owner    User
}

var (
	UserCol         = db.DB().C("users")
	AccessTokenCol  = db.DB().C("access_token")
	RefreshTokenCol = db.DB().C("refresh_token")
)
