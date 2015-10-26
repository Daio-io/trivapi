package home

import "github.com/gin-gonic/gin"

// GetHomepage - Get the homepage
func GetHomepage(c *gin.Context) {
	c.HTML(200, "home.tmpl", gin.H{
		"title": "Trivapi API",
	})
}
