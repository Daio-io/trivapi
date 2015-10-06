package main

import (
	"github.com/gin-gonic/gin"
	"trivapi/app/utils"
	"trivapi/app/trivia"
)

func main() {
	router := gin.Default()
	router.GET("/status", func(c *gin.Context) {
		c.String(200, "OK")
	})
	router.GET("/trivia", trivia.GetRandomTrivia)
	router.Run(":" + utils.GetPort())
}
