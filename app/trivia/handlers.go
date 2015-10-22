package trivia

import (
	"github.com/gin-gonic/gin"
	"trivapi/db"
)

// GetTriviaSet - Get a Random Set of Trivia handler
func GetTriviaSet(c *gin.Context) {
	options := c.MustGet(optionsKey).(db.QueryOptions)
	results := RandomTrivia(options)
	returnResponse(c, results)
}

// Private Helpers
func returnResponse(c *gin.Context, results []triviaModel) {
	if len(results) == 0 {
		c.JSON(200, gin.H{"status": "failed", "response": results})
	} else {
		c.JSON(200, gin.H{"status": "success", "response": results})
	}
}
