package http

import (
	"github.com/gin-gonic/gin"
	"moneywaste/repository"
)

type Handlers struct {
	router *gin.Engine

	userStruct *repository.User
}

func NewHandlers() *Handlers {
	db := repository.GetDB()

	return &Handlers{
		userStruct: repository.NewUser(db),
	}
}

func (h *Handlers) RunRouters() {
	h.getRoutes()

	err := h.router.Run(":8000")
	if err != nil {
		return
	}
}

func (h *Handlers) getRoutes() {
	authHandler := NewAuthHandler(h.userStruct)
	userHandler := NewUserHandler(h.userStruct)

	v1 := h.router.Group("/v1")

	userHandler.userHandlers(v1)
	authHandler.authHandlers(v1)
}
