package utils

import (
	"golang.org/x/crypto/bcrypt"
	"encoding/base64"
	"time"
)

func GenerateRefreshToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(token), nil
}

func HashRefreshToken(token string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CompareRefreshToken(hash, token string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(token))
}
