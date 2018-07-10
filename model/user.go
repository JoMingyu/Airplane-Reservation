package model

import "Airplane-Reservation/db"

type User struct {
	ID   string `bson:"_id"`
	Pw   string
	Name string
}

type Key struct {
	Owner     User
	UserAgent string
}

type Token struct {
	Key      Key
	Identity string `bson:"_id"`
}

var (
	UserCol         = db.DB().C("users")
	AccessTokenCol  = db.DB().C("access_token")
	RefreshTokenCol = db.DB().C("refresh_token")
)
