package domain

import (
	"github.com/google/uuid"
	"time"
)

type Membership struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Status    bool      `gorm:"column:status"`
	Archived  bool      `gorm:"column:archived"`
	Admin     bool      `gorm:"column:admin"`
	HidedAt   time.Time `gorm:"column:hided_at"`
	CleanedAt time.Time `gorm:"column:cleaned_at"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UserID    uuid.UUID `gorm:"column:user_id"`
	ChatID    uuid.UUID `gorm:"column:chat_id"`
}
