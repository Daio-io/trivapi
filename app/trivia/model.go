package trivia

// Model - Trivia model for trivia questions
type triviaModel struct {
	Question   string   `json:"question"`
	Answer     string   `json:"answer"`
	Options    []string `json:"options"`
	Category   string   `json:"catergory"`
	Difficulty string   `json:"difficulty"`
}
