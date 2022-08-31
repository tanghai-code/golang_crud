package main

import (
	"github.com/gin-gonic/gin"
	"golang_crud/router"
	"net/http"
)

// main() start a web service
func main() {
	// gin default server
	r := gin.Default()

	// Create a health monitor handler
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "health",
		})
	})

	// To set all router
	router.SetAllRouter(r)

	// Run the server at port 8080
	r.Run(":8080")
}
