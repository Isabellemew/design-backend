package models

import "time"

type Message struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Message   string    `json:"message" binding:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}