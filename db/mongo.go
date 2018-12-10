package db

import "github.com/globalsign/mgo"

func Connect(name string) (*mgo.Session, *mgo.Collection) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	return session, session.DB("test").C(name)
}
