package trivia

import (
	"gopkg.in/mgo.v2/bson"
	"trivapi/db"
)

// Options - Search options
type Options struct {
	Limit int
}

// QueryTrivia - Get trivia with field query
func QueryTrivia(options Options, params ...string) (*[]triviaModel, error) {
	session := db.NewSession()
	defer session.Close()
	results := []triviaModel{}
	col := session.GetTriviaCollection()

	var err error
	if params == nil {
		err = col.Find(bson.M{}).Limit(options.Limit).All(&results)
	} else {
		err = col.Find(bson.M{params[0]: params[1]}).Limit(options.Limit).All(&results)
	}
	return &results, err
}
