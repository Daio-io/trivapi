package trivia

import (
	"github.com/gin-gonic/gin"
	"trivapi/db"
)

// GetTriviaForCategory - Get all trivia for category
func GetTriviaForCategory(c *gin.Context) {
	options := c.MustGet("options").(db.QueryOptions)
	results, err := QueryTrivia(options)
	returnResponse(c, results, err)
}

// GetAllTrivia - Get all trivia
func GetAllTrivia(c *gin.Context) {
	options := c.MustGet("options").(db.QueryOptions)
	results, err := QueryTrivia(options)
	returnResponse(c, results, err)
}

// Private Helpers
func returnResponse(c *gin.Context, results interface{}, err error) {
	if err != nil {
		c.JSON(200, gin.H{"status": "failed", "response": results})
	} else {
		c.JSON(200, gin.H{"status": "success", "response": results})
	}
}
