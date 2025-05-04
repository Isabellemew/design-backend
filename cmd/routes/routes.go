package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Isabellemew/design-backend/internal/handlers"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/api/Orders", handlers.CreateOrder)
}
