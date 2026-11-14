package controller

import (
	"event-app/config"
	"event-app/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {
	var newEvent model.Event
	if err := c.ShouldBindJSON(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newEvent.DateTime.IsZero() {
		newEvent.DateTime = time.Now()
	}

	config.DB.Create(&newEvent)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   newEvent,
	})
}

func GetEvents(c *gin.Context) {
	var events []model.Event
	config.DB.Find(&events)

	if len(events) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No events found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

func GetEventByID(c *gin.Context) {
	var event model.Event
	id := c.Param("id")
	config.DB.First(&event, id)

	// Check if event exists
	err := config.DB.First(&event, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}

func UpdateEvent(c *gin.Context) {
	db := config.ConnectDatabase()
	var event model.Event
	id := c.Param("id")

	if err := db.First(&event, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&event).Updates(event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"event":   event,
	})
}

func DeleteEvent(c *gin.Context) {
	db := config.ConnectDatabase()
	var event model.Event
	id := c.Param("id")

	if err := db.First(&event, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
		})
		return
	}

	if err := db.Delete(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}
