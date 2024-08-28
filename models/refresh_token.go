package models

import "time"

type RefreshToken struct {
	Token    string
	UserID   string
	IP       string
	CreatedAt time.Time
}
