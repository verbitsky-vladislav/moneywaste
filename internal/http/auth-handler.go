package http

import (
	"github.com/gin-gonic/gin"
	"moneywaste/internal/modules/auth"
	"moneywaste/repository"
)

type AuthHandler struct {
	authService *auth.Service
}

func NewAuthHandler(r *repository.User) *AuthHandler {
	return &AuthHandler{
		authService: auth.NewAuthService(r),
	}
}

func (a *AuthHandler) authHandlers(rg *gin.RouterGroup) {
	authGroup := rg.Group("/auth")

	authGroup.POST("/sign-in", func(c *gin.Context) {
		a.authService.SignIn(c)
	})

	authGroup.POST("/sign-up", func(c *gin.Context) {
		a.authService.SignUp(c)
	})

	authGroup.POST("/refresh/:id", func(c *gin.Context) {
		a.authService.Refresh(c)
	})
}
