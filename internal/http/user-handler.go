package http

import (
	"github.com/gin-gonic/gin"
	"moneywaste/internal/common/interceptors"
	"moneywaste/internal/modules/user"
	"moneywaste/repository"
)

type UserHandler struct {
	userService *user.Service
}

func NewUserHandler(r *repository.User) *UserHandler {
	return &UserHandler{
		userService: user.NewUserService(r),
	}
}

func (u *UserHandler) userHandlers(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/my", interceptors.AuthInterceptor("all"), func(c *gin.Context) {
		u.userService.GetMy(c)
	})
	users.POST("/update", interceptors.AuthInterceptor("all"), func(c *gin.Context) {
		u.userService.Update(c)
	})

}
