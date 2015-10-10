package trivia

import "github.com/gin-gonic/gin"

// GetTriviaForCategory - Get all trivia for category
func GetTriviaForCategory(c *gin.Context) {
	category := c.Params.ByName("category")
	results, err := QueryTrivia("category", category)
	returnResponse(c, results, err)
}

// GetAllTrivia - Get all trivia
func GetAllTrivia(c *gin.Context) {
	results, err := QueryTrivia()
	returnResponse(c, results, err)
}

func returnResponse(c *gin.Context, results *[]triviaModel, err error) {
	if err != nil {
		c.JSON(200, gin.H{"status": "failed", "response": results})
	} else {
		c.JSON(200, gin.H{"status": "success", "response": results})
	}
}
