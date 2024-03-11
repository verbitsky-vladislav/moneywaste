package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

const secret string = "InR5cCI6IkpXVCJ9.eyJzdWIi"

type authService struct{}

func NewAuthService() *authService { // Исправил название метода для соответствия названию структуры
	return &authService{}
}

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub  string `json:"sub"`
	Name string `json:"name"`
	Iat  int64  `json:"iat"`
	Exp  int64  `json:"exp"`
}

func (s *authService) createSignature(message string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

func (s *authService) CreateToken(payload interface{}) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	headerJson, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerEncoded := base64.RawURLEncoding.EncodeToString(headerJson)

	// Маршалинг переданного payload
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	payloadEncoded := base64.RawURLEncoding.EncodeToString(payloadJson)

	// Обратите внимание: для создания подписи необходимо передать secret
	signature := s.createSignature(headerEncoded + "." + payloadEncoded)

	return headerEncoded + "." + payloadEncoded + "." + signature, nil
}

func (s *authService) ValidateToken(token string) (bool, Payload, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false, Payload{}, errors.New("invalid token format")
	}

	headerEncoded, payloadEncoded, signatureEncoded := parts[0], parts[1], parts[2]
	signatureCheck := s.createSignature(headerEncoded + "." + payloadEncoded) // Исправлено на вызов метода

	if signatureEncoded != signatureCheck {
		return false, Payload{}, errors.New("signature verification failed")
	}

	payloadJson, err := base64.RawURLEncoding.DecodeString(payloadEncoded)
	if err != nil {
		return false, Payload{}, errors.New("error decoding payload")
	}
	var payload Payload
	err = json.Unmarshal(payloadJson, &payload)
	if err != nil {
		return false, Payload{}, errors.New("error unmarshalling payload")
	}

	if time.Now().Unix() > payload.Exp {
		return false, Payload{}, errors.New("token is expired")
	}

	if time.Now().Unix() < payload.Iat {
		return false, Payload{}, errors.New("token issue time is in the future")
	}

	return true, payload, nil
}

func (s *authService) RefreshToken(token string) (string, error) {
	// Получаем текущее время в формате Unix timestamp
	now := time.Now().Unix()

	// Валидация токена
	valid, payload, err := s.ValidateToken(token)
	if err != nil || !valid {
		return "", err
	}

	// Установка нового времени истечения на 90 секунд в будущее от текущего момента
	payload.Exp = now + 90

	// Создание нового токена с обновленной полезной нагрузкой
	newToken, err := s.CreateToken(payload)
	return newToken, err
}
