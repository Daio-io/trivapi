package main

import (
	"github.com/gin-gonic/gin"
	"trivapi/app/trivia"
	"trivapi/app/utils"
)

func main() {
	router := gin.Default()
	router.GET("/status", func(c *gin.Context) {
		c.String(200, "OK")
	})
	router.GET("/trivia", trivia.GetRandomTrivia)
	router.Run(":" + utils.GetPort())
}
