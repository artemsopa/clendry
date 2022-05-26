package domain

import (
	"time"

	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Folder struct {
	ID types.BinaryUUID `gorm:"primary_key;default:(UUID_TO_BIN(UUID()))"`

	Title     string    `gorm:"column:title"`
	CreatedAt time.Time `gorm:"column:created_at"`

	UserID     types.BinaryUUID `gorm:"column:user_id"`
	FolderFile []FolderFile     `gorm:"foreignKey:FolderID;constraint:OnDelete:CASCADE"`
}

func (p *Folder) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	p.ID = types.BinaryUUID(id)
	return err
}
