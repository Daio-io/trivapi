package main

import (
	"github.com/gin-gonic/gin"
	"trivapi/app/trivia"
	"trivapi/app/utils"
)

func main() {
	router := gin.Default()
	router.Use(trivia.ParseTriviaParams())
	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})
	router.GET("/randomise", trivia.GetTriviaSet)
	router.GET("/category/:category", trivia.GetTriviaSet)
	router.Run(":" + utils.GetPort())
}
