package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/Isabellemew/design-backend/internal/db"
	"github.com/Isabellemew/design-backend/internal/models"
	"net/http"
)

func SendMessage(c *gin.Context) {
	var msg models.Message

	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	if err := db.DB.Create(&msg).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сохранения сообщения"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Сообщение успешно отправлено"})
}
