package domain

import "github.com/google/uuid"

type FriendRequest struct {
	Status bool      `gorm:"column:status"`
	UserID uuid.UUID `gorm:"column:user_id"`
	DefID  uuid.UUID `gorm:"column:def_id"`
}
