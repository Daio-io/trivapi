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
	router.GET("/randomise", trivia.GetAllTrivia)
	router.GET("/category/:category", trivia.GetTriviaForCategory)
	router.Run(":" + utils.GetPort())
}
