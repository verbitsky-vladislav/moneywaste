package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"moneywaste/repository"
)

type Handlers struct {
	router      *gin.Engine
	userStruct  *repository.User
	transStruct *repository.Transactions
}

func NewHandlers() *Handlers {
	db := repository.GetDB()

	return &Handlers{
		router:      gin.Default(),
		userStruct:  repository.NewUser(db),
		transStruct: repository.NewTransactions(db),
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
	transactionsHandler := NewTransactionHandler(h.transStruct)

	v1 := h.router.Group("/v1")

	userHandler.userHandlers(v1)
	fmt.Println()
	authHandler.authHandlers(v1)
	fmt.Println()
	transactionsHandler.transactionsHandlers(v1)
	fmt.Println()
}
