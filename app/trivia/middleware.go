package trivia

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"trivapi/db"
)

const maxAmount = 20

// ParseTriviaParams - Query params parsing middleware
func ParseTriviaParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		options := getQueryOptions(c)
		c.Set("options", options)
		c.Next()
	}
}

func getQueryOptions(c *gin.Context) db.QueryOptions {

	amount, _ := strconv.ParseInt(c.Query("amount"), 10, 0)
	if amount > maxAmount || amount == 0 {
		amount = maxAmount
	}
	options := db.NewQuery()
	options.Amount = int(amount)

	options.AddFilter("category", c.Params.ByName("category"))
	options.AddFilter("difficulty", c.Query("difficulty"))

	return options
}
