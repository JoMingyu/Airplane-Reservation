package db

import (
	"Airplane-Reservation/config"

	mgo "gopkg.in/mgo.v2-unstable"
)

var session *mgo.Session

func Session() *mgo.Session {
	if session == nil {
		session, _ = mgo.Dial(config.MongoURI)
		session.SetSafe(&mgo.Safe{})
	}

	return session
}

func DB() *mgo.Database {
	return Session().DB(config.DBName)
}
