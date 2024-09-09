package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = "your_secret_key"

func GetJWTSecret() string {
	return jwtSecret
}

func GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString([]byte(jwtSecret))
}
