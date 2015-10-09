package trivia

import (
	"gopkg.in/mgo.v2/bson"
	"trivapi/db"
)

// QueryTrivia - Get trivia with field query
func QueryTrivia(field string, value string) (*[]triviaModel, error) {
	session := db.NewSession()
	defer session.Close()

	results := []triviaModel{}
	col := session.GetTriviaCollection()
	err := col.Find(bson.M{field: value}).All(&results)

	return &results, err
}

// QueryAllTrivia - Get all trivia for a category
func QueryAllTrivia() (*[]triviaModel, error) {
	session := db.NewSession()
	defer session.Close()

	results := []triviaModel{}
	col := session.GetTriviaCollection()
	err := col.Find(bson.M{}).All(&results)

	return &results, err
}
