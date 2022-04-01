package domain

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Text      string    `gorm:"column:text"`
	Type      string    `gorm:"column:type"`
	Hided     bool      `gorm:"column:hided"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UserID    uuid.UUID `gorm:"column:user_id"`
	ChatID    uuid.UUID `gorm:"column:chat_id"`
}
