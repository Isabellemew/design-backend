package handlers

import (
	"log"
	"net/http"
    "github.com/Isabellemew/design-backend/internal/db"
	"github.com/Isabellemew/design-backend/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
    var user models.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        log.Println("Ошибка хэширования пароля:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать пользователя"})
        return
    }

    user.PasswordHash = string(hashedPassword)

    if err := db.DB.Create(&user).Error; err != nil {
        log.Println("Ошибка при создании пользователя:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать пользователя"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно зарегистрирован"})
}