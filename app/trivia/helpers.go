package trivia

import (
	"math/rand"
	"time"
)

type filterFunc func(triviaModel) bool

// Shuffle - shuffle array to randomise
func Shuffle(questions []triviaModel) []triviaModel {
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))

	for i := len(questions) - 1; i > 0; i-- {
		j := rand.Intn(i)
		questions[i], questions[j] = questions[j], questions[i]
	}

	return questions
}

// Filter - Filter results function
func Filter(fn filterFunc, questions []triviaModel) []triviaModel {
	outputArray := []triviaModel{}
	for i := 0; i < len(questions); i++ {
		if fn(questions[i]) {
			outputArray = append(outputArray, questions[i])
		}
	}
	return outputArray
}

// LimitArray - Splice size of array
func LimitArray(questions []triviaModel, limit int) []triviaModel {
	if len(questions) > 0 && limit <= len(questions) {
		return questions[0:limit]
	}
	return questions
}
