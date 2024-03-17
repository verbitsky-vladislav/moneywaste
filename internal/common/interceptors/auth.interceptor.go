package interceptors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"moneywaste/internal/modules/auth"
	"net/http"
)

const (
	User  string = "user"
	Admin string = "admin"
	All   string = "all"
)

func AuthInterceptor(role string) gin.HandlerFunc {
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

		if role != All {
			if role == claims.Subject {
				c.Set("id", claims.Id)
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Roles is not matched. Your role is %s, required %s", claims.Subject, role)})
				c.Abort()
				return
			}
		}

		c.Set("id", claims.Id)
		c.Next()
	}
}
