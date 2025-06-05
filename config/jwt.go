package config

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte("supersecretkey")

func GenerateJWT(userID uint, username string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(1 * time.Hour).Unix(),
	})

	return token.SignedString(JwtSecret)
}