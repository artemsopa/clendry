package domain

import (
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FolderFile struct {
	ID types.BinaryUUID `gorm:"primary_key;default:(UUID_TO_BIN(UUID()))"`

	FolderID types.BinaryUUID `gorm:"column:folder_id"`
	FileID   types.BinaryUUID `gorm:"column:file_id"`
}

func (p *FolderFile) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	p.ID = types.BinaryUUID(id)
	return err
}
