package domain

// import (
// 	"github.com/artomsopun/clendry/clendry-api/pkg/types"
// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// 	"time"
// )

// type Membership struct {
// 	ID types.BinaryUUID `gorm:"primary_key;default:(UUID_TO_BIN(UUID()))"`

// 	Archived  bool      `gorm:"column:archived"`
// 	Admin     bool      `gorm:"column:admin"`
// 	HidedAt   time.Time `gorm:"column:hided_at"`
// 	CleanedAt time.Time `gorm:"column:cleaned_at"`
// 	CreatedAt time.Time `gorm:"column:created_at"`

// 	UserID types.BinaryUUID `gorm:"column:user_id"`
// 	ChatID types.BinaryUUID `gorm:"column:chat_id"`
// }

// func (p *Membership) BeforeCreate(tx *gorm.DB) error {
// 	id, err := uuid.NewRandom()
// 	p.ID = types.BinaryUUID(id)
// 	return err
// }
