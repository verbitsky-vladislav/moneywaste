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

// SignUp - зарегистрироваться
// Закинуть в БД
// Дать куки
func (s *Service) SignUp(c *gin.Context) {
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

	jwt, err := createToken(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("jwt", jwt, 3600, "/", "localhost", false, false)
	c.JSON(http.StatusOK, id)
}

// SignIn - войти
// Валидирую юзера
// Закинуть в куки
func (s *Service) SignIn(c *gin.Context) {
	var input models.UserCreate

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := s.userRepo.GetUserByNickname(input.Nickname)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = comparePassword(user.Password, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := createToken(user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("jwt", jwt, 3600, "/", "localhost", false, false)
	c.JSON(http.StatusOK, user.Id)
}

// Refresh
// Просто выдать новый токен
func (s *Service) Refresh(c *gin.Context) {
	var id int

	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	jwt, err := createToken(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("jwt", jwt, 3600, "/", "localhost", false, false)
	c.JSON(http.StatusOK, id)
}
