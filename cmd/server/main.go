package main

import (
	"backend/internal/db"
	"backend/internal/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // уточни фронтенд-URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")
	{
		api.POST("/users", handlers.Register)
		api.POST("/login", handlers.Login)
		api.POST("/products/search", GetProducts)
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Backend работает!"})
	})

	router.OPTIONS("/*path", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Status(204)
	})

	db.InitDB() // подключение к БД
	router.Run(":8080")
}

func GetProducts(c *gin.Context) {
    query := c.Query("q")

    var products []db.Product
    if err := db.DB.Where("LOWER(name) LIKE LOWER(?)", "%"+query+"%").Find(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при запросе к базе данных"})
        return
    }

    c.JSON(http.StatusOK, products)
}
