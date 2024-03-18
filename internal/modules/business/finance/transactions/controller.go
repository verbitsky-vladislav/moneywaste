package transactions

import (
	"github.com/gin-gonic/gin"
	"moneywaste/repository"
	"moneywaste/repository/models"
	"net/http"
)

type Service struct {
	transRepo *repository.Transactions
}

func NewUserService(r *repository.Transactions) *Service {
	return &Service{
		transRepo: r,
	}
}

func (s *Service) Create(c *gin.Context) {
	var input models.Transaction // Предположим, что есть отдельная структура для входных данных регистрации

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idInterface, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized or id not found"})
		return
	}
	// Безопасное приведение типа к строке
	idStr, ok := idInterface.(string)
	if !ok {
		// Если приведение типа не удалось, значит значение не строка
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be a string"})
		return
	}

	transaction := models.Transaction{
		UserId:      idStr,
		Type:        input.Type,
		Amount:      input.Amount,
		CategoryId:  input.CategoryId,
		Date:        input.Date,
		Description: input.Description,
	}
	newTransaction, err := s.transRepo.Create(transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": newTransaction})
}

func (s *Service) Update(c *gin.Context) {

}

func (s *Service) Delete(c *gin.Context) {

}

func (s *Service) GetOneById(c *gin.Context) {

}

func (s *Service) GetAll(c *gin.Context) {

}
