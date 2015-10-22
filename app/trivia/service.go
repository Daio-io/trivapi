package trivia

import "trivapi/db"

// Temporary solution - todo: cache properly
var cachedResults []triviaModel

// RandomTrivia - Get random trivia questions
func RandomTrivia(options db.QueryOptions) []triviaModel {
	cachedTrivia := getCachedTrivia()
	shuffledResults := Shuffle(cachedTrivia)
	filterFunction := getFilterFunction(options.GetFilters())
	filteredResults := Filter(filterFunction, shuffledResults)
	return LimitArray(filteredResults, options.Amount)
}

func getCachedTrivia() []triviaModel {
	if cachedResults == nil {
		session := db.NewSession()
		defer session.Close()
		model := []triviaModel{}
		col := session.Collection(collectionName, model)
		results, _ := col.All()
		cachedResults = results.([]triviaModel)
	}
	return cachedResults
}

func getFilterFunction(filters map[string]string) filterFunc {
	cat := filters[categoryFilter]
	diff := filters[difficultyFilter]
	return func(t triviaModel) bool {
		if cat != "" && diff != "" {
			return t.Category == cat && t.Difficulty == diff
		} else if cat != "" {
			return t.Category == cat
		} else if diff != "" {
			return t.Difficulty == diff
		}
		return true
	}
}
