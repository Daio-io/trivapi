package trivia

// Difficulty consts
const (
	EASY = iota
	MEDIUM
	HARD
)

// Model - Trivia model for trivia questions
type Model struct {
	Question   string   `json:"question"`
	Answer     string   `json:"answer"`
	Choices    []string `json:"choices"`
	Difficulty int      `json:"difficulty"`
}
