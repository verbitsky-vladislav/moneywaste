package auth

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	jwt.StandardClaims
}

var tokenExpiration time.Duration = 60 * 24 * 90 * time.Minute // 90 days

func createToken(id int) (string, error) {
	expirationTime := time.Now().Add(tokenExpiration) // Токен истекает через 90 days

	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			Id:        strconv.Itoa(id),
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Валидация JWT токена
func ValidateToken(tokenString string) (*Claims, bool) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, false
	}

	if !token.Valid {
		return nil, false
	}

	// Проверяем, не истек ли срок действия токена
	if time.Now().Unix() > claims.ExpiresAt {
		return nil, false
	}

	return claims, true
}

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func comparePassword(hashedPassword, testPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))
	if err != nil {
		return false, err
	}
	// Пароли совпадают
	return true, nil
}
