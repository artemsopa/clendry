package domain

import (
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID types.BinaryUUID `gorm:"primary_key;default:(UUID_TO_BIN(UUID()))"`

	Text      string    `gorm:"column:text"`
	Type      string    `gorm:"column:type"`
	Forwarded bool      `gorm:"column:forwarded"`
	Hided     bool      `gorm:"column:hided"`
	Liked     bool      `gorm:"column:hided"`
	CreatedAt time.Time `gorm:"column:created_at"`

	UserID types.BinaryUUID `gorm:"column:user_id"`
	ChatID types.BinaryUUID `gorm:"column:chat_id"`

	Files []File `gorm:"foreignKey:MessageID;constraint:OnDelete:CASCADE"`
}

func (p *Message) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	p.ID = types.BinaryUUID(id)
	return err
}
