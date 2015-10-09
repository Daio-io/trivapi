package trivia

import "github.com/gin-gonic/gin"

// GetTriviaForCategory - Get all trivia for category
func GetTriviaForCategory(c *gin.Context) {
	category := c.Params.ByName("category")
	results, err := GetAllTrivia("category", category)
	if err != nil {
		c.JSON(200, gin.H{"status": "not found", "response": results})
	} else {
		c.JSON(200, gin.H{"status": "success", "response": results})
	}
}
