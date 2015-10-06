package trivia

import (
  "github.com/gin-gonic/gin"
  "trivapi/db"
)

// GetRandomTrivia - Get a random question
func GetRandomTrivia(c *gin.Context) {
  session := db.Connect()
  defer session.Close()

  results := []triviaModel{}
  col := session.DB("").C("trivia")
  err := col.Find(nil).All(&results)
  if err != nil {
    c.JSON(200, gin.H{"status": "not found"})
  } else {
    c.JSON(200, gin.H{"status": "success", "response": &results})
  }
}
