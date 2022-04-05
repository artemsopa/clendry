package repository

import (
	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"gorm.io/gorm"
)

type Users interface {
	GetAllWithoutBlocks(userID types.BinaryUUID) ([]domain.User, error)

	GetAllBlockedUsers(userID types.BinaryUUID) ([]domain.User, error)
	GetAllFriends(userID types.BinaryUUID) ([]domain.User, error)
	GetAllSentReqs(userID types.BinaryUUID) ([]domain.User, error)
	GetAllIncomingReqs(userID types.BinaryUUID) ([]domain.User, error)

	GetAll(userID types.BinaryUUID) ([]domain.User, error)
	GetById(userID types.BinaryUUID) (domain.User, error)
	GetByCredentials(nickname, password string) (domain.User, error)
	Create(user domain.User) error
	ChangePassword(userID types.BinaryUUID, password string) error
	Delete(userID types.BinaryUUID) error
}

type Sessions interface {
	GetRefreshToken(refreshToken string) (domain.Session, error)
	SetSession(session domain.Session) error
}

type FriendRequests interface {
	CreateUnconfirmed(request domain.FriendRequest) error
	UpdateConfirmation(request domain.FriendRequest) error
	DeleteReq(userID, defID types.BinaryUUID, status bool) error
	IsUserInFriend(userID, defID types.BinaryUUID) bool
	IsIncomingReq(userID, defID types.BinaryUUID) bool
	IsSentReq(userID, defID types.BinaryUUID) bool
}

type BlockRequests interface {
	Create(block domain.BlockRequest) error
	Delete(userID, addresseeID types.BinaryUUID) error
	IsDefInBlock(userID, defID types.BinaryUUID) bool
	IsUserInBlock(userID, defID types.BinaryUUID) bool
}

type Messages interface {
	GetLastIncomingMessagesByUserID(userID uint) ([]domain.Message, error)
	GetAddresseeMessagesByUserID(userID, addresseeID uint) ([]domain.Message, error)
	Create(message domain.Message) error
	Delete(userID, messageID uint) error
}

type Files interface {
	GetAllFilesByUserID(userID types.BinaryUUID) ([]domain.File, error)
	GetFileByUserID(userID, fileID types.BinaryUUID) (domain.File, error)

	GetAvatarByUserID(userID types.BinaryUUID) (domain.File, error)
	ChangeAvatarByUserID(userID, fileID types.BinaryUUID) error
	CreateAvatarByUserID(file domain.File) error

	GetAllTypeFilesByUserID(userID types.BinaryUUID, filetype domain.FileType) ([]domain.File, error)
	Create(file domain.File) error
	DeleteByID(userID, fileID types.BinaryUUID) error

	/*
		GetAvatarByChatID(chatID types.BinaryUUID) (domain.File, error)
		ChangeChatAvatarByMemberID(memberID, chatID types.BinaryUUID) (domain.File, error)
		CreateChatAvatarByMemberID(memberID, chatID types.BinaryUUID, file domain.File) error

		GetAllFilesByMessageID(messageID types.BinaryUUID) (domain.File, error)
		CreateFileByMessageID(chatID types.BinaryUUID) (domain.File, error)
	*/
}

type Repositories struct {
	Users          Users
	Sessions       Sessions
	Files          Files
	BlockRequests  BlockRequests
	FriendRequests FriendRequests
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users:          NewUsersRepo(db),
		Sessions:       NewSessionsRepo(db),
		Files:          NewFilesRepo(db),
		BlockRequests:  NewBlocksRepo(db),
		FriendRequests: NewFriendsRepo(db),
	}
}
