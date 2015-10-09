package trivia

import (
	"gopkg.in/mgo.v2/bson"
	"trivapi/db"
)

// GetAllTrivia - Get all trivia for a category
func GetAllTrivia(field string, value string) (*[]triviaModel, error) {
	session := db.NewSession()
	defer session.Close()

	results := []triviaModel{}
	col := session.GetTriviaCollection()
	err := col.Find(bson.M{field: value}).All(&results)

	return &results, err
}
