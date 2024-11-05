package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the home page!",
		})
	})
	r.GET("/users", func(c *gin.Context) {
		users := []string{"Alice", "Bob", "Charlie"}
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	})

	r.Run(":8080")
}
