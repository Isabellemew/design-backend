package handlers

import (
	"net/http"

	"github.com/Isabellemew/design-backend/internal/db"
	"github.com/Isabellemew/design-backend/internal/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	if err := db.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при загрузке пользователей"})
		return
	}
	c.JSON(http.StatusOK, users)
}
