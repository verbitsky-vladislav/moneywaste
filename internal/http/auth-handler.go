package http

import (
	"github.com/gin-gonic/gin"
	"moneywaste/repository"
	"net/http"
)

type AuthHandler struct {
	userRepository *repository.User
}

func NewAuthHandler(r *repository.User) *AuthHandler {
	return &AuthHandler{
		userRepository: r,
	}
}

func (a *AuthHandler) authHandlers(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")

	auth.POST("/sign-in", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users")
	})
	auth.POST("/sign-up", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users comments")
	})
	auth.POST("/refresh", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users pictures")
	})
}
