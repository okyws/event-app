package model

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	UserID      uint      `json:"user_id" binding:"-"`
	User        User      `json:"-" binding:"-" gorm:"foreignKey:UserID;references:ID"`
	DateTime    time.Time `json:"datetime"`
}
