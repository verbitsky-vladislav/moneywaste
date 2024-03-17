package user

import (
	"github.com/gin-gonic/gin"
	"moneywaste/repository"
	"net/http"
)

type Service struct {
	userRepo *repository.User
}

func NewUserService(r *repository.User) *Service {
	return &Service{
		userRepo: r,
	}
}

// GetMy - зарегистрироваться
func (s *Service) GetMy(c *gin.Context) {
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

	user, err := s.userRepo.GetUserById(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Update - войти
func (s *Service) Update(c *gin.Context) {

}

// GetAll - поулчить всех
func (s *Service) GetAll(c *gin.Context) {

}
