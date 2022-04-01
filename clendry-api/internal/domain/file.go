package domain

import (
	"github.com/google/uuid"
	"time"
)

type (
	FileStatus int
	FileType   string
)

const (
	ClientUploadInProgress FileStatus = iota
	UploadedByClient
	ClientUploadError
	StorageUploadInProgress
	UploadedToStorage
	StorageUploadError
)

const (
	Image FileType = "image"
	Video FileType = "video"
	Other FileType = "other"
)

type File struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Title       string    `gorm:"column:title"`
	Size        string    `gorm:"column:size"`
	Current     bool      `gorm:"column:current"`
	ContentType string    `gorm:"column:content_type"`
	Type        FileType  `gorm:"column:type"`
	//Status      FileStatus `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UserID    uuid.UUID `gorm:"column:user_id"`
	ChatID    uuid.UUID `gorm:"column:chat_id"`
	MessageID uuid.UUID `gorm:"column:message_id"`
}
