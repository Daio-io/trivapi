package trivia

import (
	"gopkg.in/mgo.v2/bson"
	"trivapi/db"
)

// QueryOptions - Search options
type QueryOptions struct {
	Amount int
}

// QueryTrivia - Get trivia with field query
func QueryTrivia(options QueryOptions, params ...string) (*[]triviaModel, error) {
	session := db.NewSession()
	defer session.Close()
	results := []triviaModel{}
	col := session.GetTriviaCollection()

	var err error
	if params == nil {
		err = col.Find(bson.M{}).Limit(options.Amount).All(&results)
	} else {
		err = col.Find(bson.M{params[0]: params[1]}).Limit(options.Amount).All(&results)
	}
	return &results, err
}
