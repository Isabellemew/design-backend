package handlers

import (
    "github.com/Isabellemew/design-backend/internal/db"
	"github.com/Isabellemew/design-backend/internal/models"
    "github.com/gin-gonic/gin"
    "net/http"
    "golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
    var user models.User
    var input models.User

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Login successful",
        "user": gin.H{
            "id":    user.ID,
            "name":  user.Name,
            "email": user.Email,
        },
        "token": "example_token", // временно, потом можно сделать JWT
    })
}