package domain

import (
	"github.com/google/uuid"
	"time"
)

type Chat struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Title     string    `gorm:"column:title"`
	Group     bool      `gorm:"column:group"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UserID    uuid.UUID `gorm:"column:user_id"`
	ChatID    uuid.UUID `gorm:"column:chat_id"`
}
