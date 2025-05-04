package models

import (
	"time"
)

type Order struct {
	ID          uint        `json:"id" gorm:"primaryKey"`
	UserID      uint        `json:"user_id" binding:"required"`
	User        User        `json:"-" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time   `json:"created_at" gorm:"autoCreateTime"`
	TotalAmount float64     `json:"total_amount" binding:"required" gorm:"type:decimal(10,2);not null"`
	Status      string      `json:"status" gorm:"type:varchar(50);default:'pending'"`
	Items       []OrderItem `json:"items" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
}

type OrderItem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	OrderID   uint      `json:"order_id" gorm:"index"`
	Name      string    `json:"name" binding:"required" gorm:"type:varchar(255);not null"`
	Price     float64   `json:"price" binding:"required" gorm:"type:decimal(10,2);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
