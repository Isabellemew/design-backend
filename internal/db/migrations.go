package db

import (
	"github.com/Isabellemew/design-backend/internal/models"
)

func AutoMigrate() {
	// ... existing code ...
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.OrderItem{})
	// ... existing code ...
}
