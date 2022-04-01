package domain

import (
	"github.com/google/uuid"
)

type BlockRequest struct {
	UserID uuid.UUID `gorm:"column:user_id"`
	DefID  uuid.UUID `gorm:"column:def_id"`
}
