package main

import (
	"github.com/gin-gonic/gin"
	"trivapi/app/config"
	"trivapi/app/home"
	"trivapi/app/trivia"
)

func main() {
	router := gin.Default()

	// Middleware
	router.Use(trivia.ParseTriviaParams())

	// Website / Homepage
	router.GET("/", home.GetHomepage)
	router.LoadHTMLGlob("app/templates/*")

	// API
	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})
	router.GET("/randomise", trivia.GetTriviaSet)
	router.GET("/category/:category", trivia.GetTriviaSet)
	router.Run(":" + utils.GetPort())
}
