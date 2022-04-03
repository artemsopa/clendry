package domain

import (
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID types.BinaryUUID `gorm:"primary_key;default:(UUID_TO_BIN(UUID()))"`

	Nick     string `gorm:"column:nick"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`

	Session        Session         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	SentReqs       []FriendRequest `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	IncomingReqs   []FriendRequest `gorm:"foreignKey:DefID;constraint:OnDelete:CASCADE"`
	SentBlocks     []BlockRequest  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	IncomingBlocks []BlockRequest  `gorm:"foreignKey:DefID;constraint:OnDelete:CASCADE"`
	Messages       []Message       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Memberships    []Membership    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Files          []File          `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	u.ID = types.BinaryUUID(id)
	return err
}
