package http

import (
	"github.com/gin-gonic/gin"
	"moneywaste/internal/common/interceptors"
	"moneywaste/internal/modules/business/finance/transactions"
	"moneywaste/repository"
)

type TransactionHandler struct {
	transService *transactions.Service
}

func NewTransactionHandler(r *repository.Transactions) *TransactionHandler {
	return &TransactionHandler{
		transService: transactions.NewUserService(r),
	}
}

func (u *TransactionHandler) transactionsHandlers(rg *gin.RouterGroup) {
	trans := rg.Group("/transactions")

	trans.POST("/create", interceptors.AuthInterceptor("all"), func(c *gin.Context) {
		u.transService.Create(c)
	})
	trans.DELETE("/delete", interceptors.AuthInterceptor("all"), func(c *gin.Context) {
		u.transService.Delete(c)
	})
	trans.PUT("/update", interceptors.AuthInterceptor("all"), func(c *gin.Context) {
		u.transService.Update(c)
	})
	trans.GET("/one/:id", interceptors.AuthInterceptor("all"), func(c *gin.Context) {
		u.transService.GetOneById(c)
	})
	trans.GET("/all", interceptors.AuthInterceptor("all"), func(c *gin.Context) {
		u.transService.GetAll(c)
	})

}
