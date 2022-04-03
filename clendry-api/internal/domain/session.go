package domain

import (
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Session struct {
	ID types.BinaryUUID `gorm:"primary_key;default:(UUID_TO_BIN(UUID()))"`

	RefreshToken string    `gorm:"column:refresh_token"`
	ExpiresAt    time.Time `gorm:"column:expires_at"`

	UserID types.BinaryUUID `gorm:"column:user_id;"`
}

func (s *Session) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	s.ID = types.BinaryUUID(id)
	return err
}
