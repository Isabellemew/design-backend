package handlers 

import (
	"backend/internal/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProducts(c *gin.Context) {
    query := c.Query("q")

    var products []db.Product
    if err := db.DB.Where("LOWER(name) LIKE LOWER(?)", "%"+query+"%").Find(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при запросе к базе данных"})
        return
    }

    c.JSON(http.StatusOK, products)
}
