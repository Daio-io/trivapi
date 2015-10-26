package main

import (
	"github.com/gin-gonic/gin"
	"trivapi/app/config"
	"trivapi/app/home"
	"trivapi/app/trivia"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("app/templates/*")
	router.Use(trivia.ParseTriviaParams())
	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})
	router.GET("/", home.GetHomepage)
	router.GET("/randomise", trivia.GetTriviaSet)
	router.GET("/category/:category", trivia.GetTriviaSet)
	router.Run(":" + utils.GetPort())
}
