package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"studygroup/db/postgres"
)

func main() {

	db := postgres.GetConnection()
	defer db.Close()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	r.Run("0.0.0.0:8080")
}
