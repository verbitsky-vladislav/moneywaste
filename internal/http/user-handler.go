package http

import (
	"github.com/gin-gonic/gin"
	"moneywaste/repository"
	"net/http"
)

type UserHandler struct {
	userRepository *repository.User
}

func NewUserHandler(r *repository.User) *UserHandler {
	return &UserHandler{
		userRepository: r,
	}
}

func (h *UserHandler) userHandlers(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/my", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users")
	})
	users.POST("/update", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users comments")
	})

}
