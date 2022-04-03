package repository

import (
	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"gorm.io/gorm"
)

type Users interface {
	GetAll() ([]domain.User, error)
	GetById(userID uint) (domain.User, error)
	GetByCredentials(nickname, password string) (domain.User, error)
	Create(user domain.User) error
	ChangePassword(userID uint, password string) error
	ChangeAvatar(userID uint, url string) error
	Delete(userID uint) error
}

type Sessions interface {
	GetRefreshToken(refreshToken string) (domain.Session, error)
	SetSession(session domain.Session) error
}

type FriendRequests interface {
	GetAllIncomingUnconfirmedByUserID(userID uint) ([]domain.FriendRequest, error)
	GetAllSentUnconfirmedByUserID(userID uint) ([]domain.FriendRequest, error)
	GetAllConfirmedByUserID(userID uint) ([]domain.FriendRequest, error)
	CreateUnconfirmed(request domain.FriendRequest) error
	UpdateConfirmation(request domain.FriendRequest) error
	DeleteRequest(userID, addresseeID uint) error
}

type Blocks interface {
	GetAllByUserID(userID uint) ([]domain.BlockRequest, error)
	Create(block domain.BlockRequest) error
	Delete(userID, addresseeID uint) error
}

type Messages interface {
	GetLastIncomingMessagesByUserID(userID uint) ([]domain.Message, error)
	GetAddresseeMessagesByUserID(userID, addresseeID uint) ([]domain.Message, error)
	Create(message domain.Message) error
	Delete(userID, messageID uint) error
}

type Repositories struct {
	Users    Users
	Sessions Sessions
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users:    NewUsersRepo(db),
		Sessions: NewSessionsRepo(db),
	}
}
