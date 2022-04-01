package domain

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	RefreshToken string    `gorm:"column:refresh_token"`
	ExpiresAt    time.Time `gorm:"column:expires_at"`
	UserID       uuid.UUID `gorm:"column:user_id"`
}
