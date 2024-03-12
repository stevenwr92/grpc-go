package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golangtest.com/models"
)

func GenerateJWTToken(payload models.User) (string, error) {
	claims := jwt.MapClaims{
		"email":     payload.Email,
		"createdAt": payload.CreatedAt,
		"exp":       time.Now().Add(360 * time.Second).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("Rahasia"))
}
