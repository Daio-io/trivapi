package trivia

import (
	"gopkg.in/mgo.v2/bson"
	"trivapi/db"
)

// GetAllTriviaForCategory - Get all trivia for a category
func GetAllTriviaForCategory(category string) (*[]triviaModel, error) {
	session := db.NewSession()
	defer session.Close()

	results := []triviaModel{}
	col := session.GetTriviaCollection()
	err := col.Find(bson.M{"category": category}).All(&results)

	return &results, err
}
