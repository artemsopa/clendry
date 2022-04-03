package service

import (
	"context"
	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"github.com/artomsopun/clendry/clendry-api/internal/repository"
	"github.com/artomsopun/clendry/clendry-api/pkg/auth"
	"github.com/artomsopun/clendry/clendry-api/pkg/files"
	"github.com/artomsopun/clendry/clendry-api/pkg/hash"
	"github.com/artomsopun/clendry/clendry-api/pkg/storage"
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"time"
)

type UserInfo struct {
	ID     types.BinaryUUID
	Nick   string
	Email  string
	Avatar string
}

type UserInputSigUp struct {
	Nick      string
	Email     string
	Passwords Passwords
}

type UserInputSigIn struct {
	Login    string
	Password string
}

type PasswordConfirm struct {
	UserID      types.BinaryUUID
	OldPassword string
	Passwords   Passwords
}

type Passwords struct {
	Password string
	Confirm  string
}

type Tokens struct {
	AccessToken  Token
	RefreshToken Token
}

type Token struct {
	Value     string
	ExpiresAt time.Time
}

type File struct {
	Title       string
	Size        int64
	ContentType string
	Type        domain.FileType
	CreatedAt   time.Time
	ForeignID   types.BinaryUUID
}

const (
	users    = "users"
	chats    = "chats"
	messages = "messages"
)

type Message struct {
	ID        types.BinaryUUID
	Title     string
	Text      string
	Type      string
	Holder    bool
	Forwarded bool
	CreatedAt time.Time
	File      []File
}

type Chat struct {
	ID        types.BinaryUUID
	Title     string
	Avatar    string
	Group     bool
	Cleaned   bool
	CreatedAt time.Time
}

type ChatListItem struct {
	Chat    Chat
	Message Message
}

type ChatItem struct {
	Chat     Chat
	Messages []Message
	Members  []UserInfo
}

type ChatInput struct {
	Chat    Chat
	Members []types.BinaryUUID
}

type Auths interface {
	SignUp(input UserInputSigUp) error
	SignIn(input UserInputSigIn) (Tokens, error)
	RefreshTokens(refresh Token) (Tokens, error)
}

type Profiles interface {
	GetProfile(userID types.BinaryUUID) (UserInfo, error)
	ChangePassword(confirm PasswordConfirm) error
	DeleteProfile(userID types.BinaryUUID) error

	UploadAvatar(ctx context.Context, file File) error
	ChangeAvatarByFileID(userID, fileID types.BinaryUUID) error
}

type Users interface {
	GetAllUsers(userID types.BinaryUUID) ([]UserInfo, error)        //without blocks
	GetUserByID(userID, defID types.BinaryUUID) ([]UserInfo, error) //without blocks

	GetAllFriends(userID types.BinaryUUID) ([]UserInfo, error) //status true
	GetAllIncomingRequests(userID types.BinaryUUID) ([]UserInfo, error)
	GetAllSentRequests(userID types.BinaryUUID) ([]UserInfo, error)
	SendFriendRequest(userID, defID types.BinaryUUID) error
	ConfirmFriendRequest(userID, defID types.BinaryUUID) error
	DeleteIncomingRequest(userID, defID types.BinaryUUID) error
	DeleteSentRequest(userID, defID types.BinaryUUID) error

	GetAllBlockedUsers(userID types.BinaryUUID) ([]UserInfo, error)
	CreateBlock(userID, defID types.BinaryUUID) error
	DeleteBlock(userID, defID types.BinaryUUID) error
}

type FileStorages interface {
	GetAllStorageFiles(userID types.BinaryUUID) ([]File, error)

	GetAllImageFiles(userID types.BinaryUUID) ([]File, error)
	GetAllVideoFiles(userID types.BinaryUUID) ([]File, error)
	GetAllOtherFiles(userID types.BinaryUUID) ([]File, error)

	GetStorageFile(userID, fileID types.BinaryUUID) (File, error)

	UploadImageFile(userID, fileID types.BinaryUUID) error
	UploadVideoFile(userID, fileID types.BinaryUUID) error
	UploadOtherFile(userID, fileID types.BinaryUUID) error

	DeleteStorageFile(userID, fileID types.BinaryUUID) error
}

type Chats interface {
	GetAllChats(userID types.BinaryUUID) ([]ChatListItem, error) //without hided //By membership
	GetAllArchivedChats(userID types.BinaryUUID) ([]ChatListItem, error)

	ArchiveChat(userID, chatID types.BinaryUUID) error
	RemoveChatFromArchive(userID, chatID types.BinaryUUID) error

	GetChat(userID, chatID types.BinaryUUID) (ChatItem, error)
	CreateChat(userID, defID types.BinaryUUID) error //non-admin

	AddChatMembers(userID types.BinaryUUID, defIDs []types.BinaryUUID) error

	RemoveChatMember(userID, defIDs types.BinaryUUID) error
	DeleteChat(userID, chatID types.BinaryUUID) error

	HideChat(userID, chatID types.BinaryUUID) error
	CleanChatHistory(userID, chatID types.BinaryUUID)

	CreateGroupChat(userID types.BinaryUUID, chat Chat, file File) error //with admin

	CreateGroupChatAvatar(userID types.BinaryUUID, chat Chat, file File) error //By files
	UpdateGroupChatAvatar(userID, chatID types.BinaryUUID, file File) error

	UpdateGroupChatTitle(userID, chatID types.BinaryUUID, title string) error

	SendMessage(userID types.BinaryUUID, message Message) error
	ForwardMessage(userID types.BinaryUUID, message Message) error
	DeleteMessage(userID, messageID types.BinaryUUID) error

	LeaveChatGroup(userID, chatID types.BinaryUUID) error //non-group chat can be hided only
}

type Services struct {
	Auths    Auths
	Profiles Profiles
}

type Deps struct {
	Repos           *repository.Repositories
	Hasher          hash.PasswordHasher
	FilesManager    *files.FilesManager
	TokenManager    auth.TokenManager
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	StorageProvider storage.Provider
}

func NewServices(deps Deps) *Services {
	authsService := NewAuthsService(deps.Repos.Users, deps.Repos.Sessions, deps.Hasher, deps.TokenManager, deps.AccessTokenTTL, deps.RefreshTokenTTL)
	profilesService := NewProfilesService(deps.Repos.Users, deps.Repos.Files, deps.FilesManager, deps.Hasher)

	return &Services{
		Auths:    authsService,
		Profiles: profilesService,
	}
}
