package auth

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"moneywaste/repository"
	"moneywaste/repository/models"
	"net/http"
)

type Service struct {
	userRepo *repository.User
}

func NewAuthService(r *repository.User) *Service {
	user := models.User{
		Fio:      "admin",
		Email:    "admin@gmail.com",
		Password: "admin",
	}
	id, err := r.Create(user)
	if err != nil {
		log.Fatal("admin user was not created", err)
	}

	jwt, err := createToken(id, "admin")
	if err != nil {
		log.Fatal("admin user was not created", err)
	}
	// Создаем структуру для логирования
	logData := struct {
		ID  string `json:"id"`
		JWT string `json:"jwt"`
	}{
		ID:  id,
		JWT: jwt,
	}

	// Конвертируем структуру в JSON с отступами
	data, err := json.MarshalIndent(logData, "", "    ") // используем 4 пробела для отступа
	if err != nil {
		log.Fatalf("Error marshalling log data: %v", err)
	}

	// Выводим JSON строку в лог
	log.Println(string(data))

	return &Service{
		userRepo: r,
	}
}

// SignUp - зарегистрироваться
// Закинуть в БД
// Дать куки
func (s *Service) SignUp(c *gin.Context) {
	var input models.User // Предположим, что есть отдельная структура для входных данных регистрации

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newPass, err := hashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		Fio:      input.Fio,
		Email:    input.Email,
		Password: newPass,
	}
	id, err := s.userRepo.Create(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := createToken(id, "user")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("jwt", jwt, 3600, "/", "localhost", false, false)
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// SignIn - войти
// Валидирую юзера
// Закинуть в куки
func (s *Service) SignIn(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := s.userRepo.GetOneByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = comparePassword(user.Password, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := createToken(user.Id, "user")
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
	var id string

	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	jwt, err := createToken(id, "user")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("jwt", jwt, 3600, "/", "localhost", false, false)
	c.JSON(http.StatusOK, id)
}
