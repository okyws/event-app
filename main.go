package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	r.GET("/test", getTestHandler)

	// Start server on port 8080 (default)
	r.Run(":8080")
}

func getTestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
		"data":    "isi test data",
	})
}
