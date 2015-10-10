package trivia

import (
	"gopkg.in/mgo.v2/bson"
	"trivapi/db"
)

// QueryTrivia - Get trivia with field query
func QueryTrivia(params ...string) (*[]triviaModel, error) {
	session := db.NewSession()
	defer session.Close()
	results := []triviaModel{}
	col := session.GetTriviaCollection()

	var err error
	if params == nil {
		err = col.Find(bson.M{}).All(&results)
	} else {
		err = col.Find(bson.M{params[0]: params[1]}).All(&results)
	}
	return &results, err
}
