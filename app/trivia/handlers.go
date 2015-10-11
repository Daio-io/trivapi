package trivia

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

const defaultAmount = 20

// GetTriviaForCategory - Get all trivia for category
func GetTriviaForCategory(c *gin.Context) {
	options := getQueryOptions(c)
	category := c.Params.ByName("category")
	results, err := QueryTrivia(options, "category", category)
	returnResponse(c, results, err)
}

// GetAllTrivia - Get all trivia
func GetAllTrivia(c *gin.Context) {
	options := getQueryOptions(c)
	results, err := QueryTrivia(options)
	returnResponse(c, results, err)
}

func returnResponse(c *gin.Context, results *[]triviaModel, err error) {
	if err != nil {
		c.JSON(200, gin.H{"status": "failed", "response": results})
	} else {
		c.JSON(200, gin.H{"status": "success", "response": results})
	}
}

func getQueryOptions(c *gin.Context) QueryOptions {
	amount, err := strconv.ParseInt(c.Query("amount"), 10, 0)
	if err != nil || amount > defaultAmount {
		amount = defaultAmount
	}
	options := QueryOptions{Amount: int(amount)}
	return options
}
