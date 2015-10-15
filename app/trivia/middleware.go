package trivia

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"trivapi/db"
)

const (
	defaultAmount = 20
)

// ParseTriviaParams - Get all trivia for category
func ParseTriviaParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		options := getQueryOptions(c)
		c.Set("options", options)
		c.Next()
	}
}

func getQueryOptions(c *gin.Context) db.QueryOptions {

	amount, _ := strconv.ParseInt(c.Query("amount"), 10, 0)
	if amount > defaultAmount {
		amount = defaultAmount
	}
	options := db.NewQuery()
	options.Amount = int(amount)

	options.AddFilter("category", c.Params.ByName("category"))
	options.AddFilter("difficulty", c.Query("difficulty"))

	return options
}
