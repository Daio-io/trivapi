package db

import "gopkg.in/mgo.v2"

var connectionString = getConnectionString()

// Connect to database - returns a session
func Connect() *mgo.Session {
	session, err := mgo.Dial(connectionString)
	if err != nil {
		panic(err)
	}
	return session
}
