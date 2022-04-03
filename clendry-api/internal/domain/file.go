package domain

import (
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type FileType string

const (
	Image FileType = "image"
	Video FileType = "video"
	Other FileType = "other"
)

type File struct {
	ID types.BinaryUUID `gorm:"primary_key;default:(UUID_TO_BIN(UUID()))"`

	Title       string    `gorm:"column:title"`
	Size        string    `gorm:"column:size"`
	Current     bool      `gorm:"column:current"`
	ContentType string    `gorm:"column:content_type"`
	Type        FileType  `gorm:"column:type"`
	CreatedAt   time.Time `gorm:"column:created_at"`

	UserID    types.BinaryUUID `gorm:"column:user_id"`
	ChatID    types.BinaryUUID `gorm:"column:chat_id"`
	MessageID types.BinaryUUID `gorm:"column:message_id"`
}

func (p *File) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	p.ID = types.BinaryUUID(id)
	return err
}
