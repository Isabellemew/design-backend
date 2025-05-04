package main

import (
	"time"

	"github.com/Isabellemew/design-backend/internal/db"
	"github.com/Isabellemew/design-backend/internal/handlers"
	"github.com/Isabellemew/design-backend/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Статические файлы
	router.Static("/products", "./products")

	// ✅ Подключаем middleware для базы данных — ДО маршрутов
	router.Use(middleware.DBMiddleware())

	// Роуты API
	api := router.Group("/api")
	{
		api.POST("/users", handlers.Register)
		api.POST("/login", handlers.Login)
		api.GET("/products/search", handlers.SearchProducts)
		api.GET("/products", handlers.GetProducts)
		api.GET("/categories", handlers.GetCategories)
		api.POST("/Message", handlers.SendMessage)
		api.POST("/orders", handlers.CreateOrder)
		api.GET("/orders", handlers.GetOrders)
		api.GET("/orders/:id", handlers.GetOrder)
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Backend работает!"})
	})

	// CORS preflight
	router.OPTIONS("/*path", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Status(204)
	})

	// Инициализация БД
	db.InitDB()

	// Запуск сервера
	router.Run(":8080")
}
