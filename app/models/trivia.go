package models

// Difficulty consts
const (
	EASY = iota
	MEDIUM
	HARD
)

// TriviaModel - Model for Triva Questions
type TriviaModel struct {
	Question   string
	Answer     string
	Choices    []string
	Difficulty int
}
