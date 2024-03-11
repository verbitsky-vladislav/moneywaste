package auth

import (
	"github.com/gin-gonic/gin"
	"moneywaste/repository"
	"moneywaste/repository/models"
	"net/http"
)

type Service struct {
	userRepo *repository.User
}

func NewAuthService(r *repository.User) *Service {
	return &Service{
		userRepo: r,
	}
}

func (s Service) SignUp(c *gin.Context) {
	var input models.UserCreate

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newPass, err := hashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.UserCreate{Nickname: input.Nickname, Password: newPass}
	id, err := s.userRepo.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, id)
}

func (s Service) SignIn(c *gin.Context) {

}

func (s Service) Refresh(c *gin.Context) {

}
