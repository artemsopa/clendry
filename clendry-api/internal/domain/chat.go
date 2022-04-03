package domain

import (
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Chat struct {
	ID types.BinaryUUID `gorm:"primary_key;default:(UUID_TO_BIN(UUID()))"`

	Title     string    `gorm:"column:title"`
	Group     bool      `gorm:"column:group"`
	CreatedAt time.Time `gorm:"column:created_at"`

	Messages    []Message    `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE"`
	Memberships []Membership `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE"`
	Files       []File       `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE"`
}

func (p *Chat) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	p.ID = types.BinaryUUID(id)
	return err
}
