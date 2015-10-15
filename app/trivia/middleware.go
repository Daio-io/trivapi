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
	filters := map[string]string{}

	filters = addFilter("category", c.Params.ByName("category"), filters)
	filters = addFilter("difficulty", c.Query("difficulty"), filters)

	options := db.QueryOptions{Amount: int(amount), Filters: filters}
	return options
}

// Refactor into filter object with helpers
func addFilter(filterName string, value string, filters map[string]string) map[string]string {
	if value != "" {
		filters[filterName] = value
	}
	return filters
}
