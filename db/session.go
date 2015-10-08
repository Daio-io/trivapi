package db

import "gopkg.in/mgo.v2"

const collectionName = "trivia"

// Session - Session Struct
type Session struct {
	session *mgo.Session
}

// NewSession - sessions
func NewSession() *Session {
	ses := new(Session)
	ses.session = connect().Copy()
	return ses
}

// Close - close the session
func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}

// GetTriviaCollection - get the collection
func (s *Session) GetTriviaCollection() *mgo.Collection {
	if s.session != nil {
		return s.session.DB("").C(collectionName)
	}
	panic("Attempt to get Trivia Collection when session is nil")
}
