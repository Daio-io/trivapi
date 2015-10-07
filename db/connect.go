package db

import "gopkg.in/mgo.v2"

var connectionString = getConnectionString()
var session, err = mgo.Dial(connectionString)

// Connect to database - returns a session
func Connect() *mgo.Session {
	if err != nil {
		panic(err)
	}
	return session.Copy()
}
