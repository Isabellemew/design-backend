package handlers 

import (
	"github.com/gin-gonic/gin"
	"net/http"
    "github.com/Isabellemew/design-backend/internal/db"
	"github.com/Isabellemew/design-backend/internal/models"
)

func SearchProducts(c *gin.Context) {
    query := c.Query("q")

    var products []models.Product
    if err := db.DB.Where("LOWER(name) LIKE LOWER(?)", "%"+query+"%").Find(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при запросе к базе данных"})
        return
    }

    c.JSON(http.StatusOK, products)
}

func GetProducts(c *gin.Context){

    var categories []models.Category
    if err := db.DB.Preload("Products").Find(&categories).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при запросе к базе данных"})
        return
    }
    // func GetAll(db *gorm.DB) ([]User, error) {
    //     var users []User
    //     err := db.Model(&User{}).Preload("CreditCards").Find(&users).Error
        // return users, err
    // }

    c.JSON(http.StatusOK, categories)
}
