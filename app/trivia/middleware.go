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
		c.Set(optionsKey, options)
		c.Next()
	}
}

func getQueryOptions(c *gin.Context) db.QueryOptions {

	amount, _ := strconv.ParseInt(c.Query(amountFilter), 10, 0)
	if amount > maxAmount || amount <= 0 {
		amount = maxAmount
	}
	options := db.NewQuery()
	options.Amount = int(amount)

	options.AddFilter(categoryFilter, c.Params.ByName(categoryFilter))
	options.AddFilter(difficultyFilter, c.Query(difficultyFilter))

	return options
}
