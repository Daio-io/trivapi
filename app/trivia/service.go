package trivia

import "trivapi/db"

const collectionName = "trivia"

var collectionCount string

// QueryTrivia - Get trivia with field query
func QueryTrivia(options db.QueryOptions) (interface{}, error) {
	session := db.NewSession()
	defer session.Close()
	model := []triviaModel{}
	col := session.Collection(collectionName, model)

	results, err := col.FindRandom(options)
	return results, err
}
