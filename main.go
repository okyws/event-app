package main

import (
	"event-app/config"
	"event-app/controller"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	db := config.ConnectDatabase()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error getting underlying database connection: ", err)
	}
	defer sqlDB.Close()

	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.GET("/events", controller.GetEvents)
		api.POST("/event", controller.CreateEvent)
		api.GET("/event/:id", controller.GetEventByID)
		api.PUT("/event/:id", controller.UpdateEvent)
		api.DELETE("/event/:id", controller.DeleteEvent)
	}

	// Start server on port 8080 (default)
	r.Run(":8080")
}
