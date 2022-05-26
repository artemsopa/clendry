package domain

import (
	"time"

	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	UrlTitle    string    `gorm:"column:url_title"`
	Size        int64     `gorm:"column:size"`
	ContentType string    `gorm:"column:content_type"`
	Type        FileType  `gorm:"column:type"`
	IsFavourite bool      `gorm:"column:is_favourite"`
	IsTrash     bool      `gorm:"column:is_trash"`
	CreatedAt   time.Time `gorm:"column:created_at"`

	UserID     types.BinaryUUID `gorm:"column:user_id"`
	FolderFile []FolderFile     `gorm:"foreignKey:FileID;constraint:OnDelete:CASCADE"`

	//ChatID    types.BinaryUUID `gorm:"column:chat_id"`
	//MessageID types.BinaryUUID `gorm:"column:message_id"`
}

func (p *File) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	p.ID = types.BinaryUUID(id)
	return err
}
