package utils

import (
	"fmt"
	"time"

	"github.com/Taala2/auth-service/config"
	"github.com/golang-jwt/jwt"
)

func GenerateAccessToken(userID, ip string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"ip":      ip,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(config.JWTSecret))
}

func ValidateAccessToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("errr")
		}
		return []byte(config.JWTSecret), nil
	})
}
