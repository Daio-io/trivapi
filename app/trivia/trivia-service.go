package trivia

import (
	"gopkg.in/mgo.v2/bson"
	"trivapi/db"
)

// GetAllTriviaForCategory - Get all trivia for a category
func GetAllTriviaForCategory(category string) (*[]triviaModel, error) {
	session := db.Connect()
	defer session.Close()

	results := []triviaModel{}
	col := session.DB("").C("trivia")
	err := col.Find(bson.M{"category": category}).All(&results)

	return &results, err
}
