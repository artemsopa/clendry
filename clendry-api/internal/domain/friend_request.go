package domain

import (
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FriendRequest struct {
	ID types.BinaryUUID `gorm:"primary_key;default:(UUID_TO_BIN(UUID()))"`

	Status bool `gorm:"column:status"`

	UserID types.BinaryUUID `gorm:"column:user_id"`
	DefID  types.BinaryUUID `gorm:"column:def_id"`
}

func (p *FriendRequest) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	p.ID = types.BinaryUUID(id)
	return err
}
