package trivia

// Difficulty consts
const (
	EASY = iota
	MEDIUM
	HARD
)

// Model - Trivia model for triva questions
type Model struct {
	Question   string
	Answer     string
	Choices    []string
	Difficulty int
}
