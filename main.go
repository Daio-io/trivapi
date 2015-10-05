package main

import (
	"github.com/gin-gonic/gin"
	"trivapi/app/utils"
)

func main() {
	router := gin.Default()
	router.GET("/status", func(c *gin.Context) {
		c.String(200, "OK")
	})
	router.Run(":" + utils.GetPort())
}
