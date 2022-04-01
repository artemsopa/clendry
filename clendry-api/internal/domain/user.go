package domain

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Nick     string    `gorm:"column:nick"`
	Email    string    `gorm:"column:email"`
	Password string    `gorm:"column:password"`

	Session Session `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	/*Friends  []Friend  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Blocks   []Block   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Messages []Message `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	*/
}
