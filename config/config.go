package config

import (
	"os"
)

var (
	JWTSecret     = os.Getenv("JWT_SECRET")
	DatabaseURL   = os.Getenv("DATABASE_URL")
	EmailSender   = "bayirbekutaalay@icloud.com" // Используйте свою почту для отправки предупреждений
)
