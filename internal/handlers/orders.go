package handlers

import (
	"net/http"
	"strconv"

	"github.com/Isabellemew/design-backend/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateOrder handles POST /api/orders
func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных", "details": err.Error()})
		return
	}

	// Get database instance
	database := c.MustGet("db").(*gorm.DB)

	// Start transaction
	tx := database.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании заказа"})
		return
	}

	// Create order with items
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении заказа", "details": err.Error()})
		return
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении заказа"})
		return
	}

	// Return created order
	c.JSON(http.StatusCreated, order)
}

// GetOrders handles GET /api/orders
func GetOrders(c *gin.Context) {
	var orders []models.Order
	database := c.MustGet("db").(*gorm.DB)

	if err := database.Preload("Items").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении заказов"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// GetOrder handles GET /api/orders/:id
func GetOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID заказа"})
		return
	}

	var order models.Order
	database := c.MustGet("db").(*gorm.DB)

	if err := database.Preload("Items").First(&order, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении заказа"})
		return
	}

	c.JSON(http.StatusOK, order)
}
