package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/Isabellemew/design-backend/internal/db"
)

func DBMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db.DB)
		c.Next()
	}
}