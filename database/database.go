package database

import (
	"gopkg.in/mgo.v2"
)

func NewPastasConnection() (*mgo.Session, *mgo.Collection, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, nil, err
	}

	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("tonnarello").C("pastas")

	return session, collection, err
}
