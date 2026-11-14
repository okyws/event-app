package main

import (
	"event-app/config"
	"event-app/controller"
	"event-app/middlewares"
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
	config.DB = config.ConnectDatabase()
	sqlDB, err := config.DB.DB()
	if err != nil {
		log.Fatal("Error getting underlying database connection: ", err)
	}
	defer sqlDB.Close()

	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("auth/register", controller.RegisterUser)
		api.POST("auth/login", controller.LoginUser)
		api.GET("/events", controller.GetEvents)
		api.GET("/event/:id", controller.GetEventByID)

		protected := api.Group("/")
		protected.Use(middlewares.AuthMiddleware())
		{
			protected.GET("/auth/me", controller.GetUser)
			protected.POST("/event", controller.CreateEvent)
			protected.PUT("/event/:id", controller.UpdateEvent)
			protected.DELETE("/event/:id", controller.DeleteEvent)
		}
	}

	// Start server on port 8080 (default)
	r.Run(":8080")
}
