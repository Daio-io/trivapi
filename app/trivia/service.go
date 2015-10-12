package trivia

import (
	"trivapi/db"
)

const collectionName = "trivia"

// QueryOptions - Search options
type QueryOptions struct {
	Amount  int
	Filters map[string]string
}

// QueryTrivia - Get trivia with field query
func QueryTrivia(options QueryOptions) (*[]triviaModel, error) {
	session := db.NewSession()
	defer session.Close()
	results := []triviaModel{}
	col := session.Collection(collectionName)

	err := col.Find(options.Filters).Limit(options.Amount).All(&results)
	return &results, err
}
