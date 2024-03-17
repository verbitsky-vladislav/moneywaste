package interceptors

import (
	"github.com/gin-gonic/gin"
	"moneywaste/internal/controllers/auth"
	"net/http"
)

func AuthInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// Проверяем, что токен присутствует
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		claims, ok := auth.ValidateToken(tokenString)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("id", claims.Id)
		c.Next()
	}
}
