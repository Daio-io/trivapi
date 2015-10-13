package trivia

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"trivapi/db"
)

const (
	defaultAmount = 20
	cat           = "category"
)

// GetTriviaForCategory - Get all trivia for category
func GetTriviaForCategory(c *gin.Context) {
	options := getQueryOptions(c)
	filters := map[string]string{cat: c.Params.ByName(cat)}
	options.Filters = filters
	results, err := QueryTrivia(options)
	returnResponse(c, results, err)
}

// GetAllTrivia - Get all trivia
func GetAllTrivia(c *gin.Context) {
	options := getQueryOptions(c)
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

func getQueryOptions(c *gin.Context) db.QueryOptions {
	amount, err := strconv.ParseInt(c.Query("amount"), 10, 0)
	if err != nil || amount > defaultAmount {
		amount = defaultAmount
	}
	options := db.QueryOptions{Amount: int(amount)}
	return options
}
