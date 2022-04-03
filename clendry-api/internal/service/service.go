package service

import (
	"context"
	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"github.com/artomsopun/clendry/clendry-api/internal/repository"
	"github.com/artomsopun/clendry/clendry-api/pkg/auth"
	"github.com/artomsopun/clendry/clendry-api/pkg/hash"
	"github.com/artomsopun/clendry/clendry-api/pkg/storage"
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"github.com/google/uuid"
	"time"
)

type UserInfo struct {
	ID     uuid.UUID
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
	Size        string
	Current     bool
	ContentType string
	Type        domain.FileType
	CreatedAt   time.Time
	ForeignID   types.BinaryUUID
}

type Message struct {
	ID        uuid.UUID
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

	GetAvatar(userID types.BinaryUUID) (UserInfo, error)
	UploadAvatar(ctx context.Context, file File) error
	ChangeAvatarByFileID(ctx context.Context, fileID types.BinaryUUID) error
	DeleteAvatar(ctx context.Context, userID types.BinaryUUID) error
}

type Users interface {
	GetAllUsers(userID uuid.UUID) ([]UserInfo, error)        //without blocks
	GetUserByID(userID, defID uuid.UUID) ([]UserInfo, error) //without blocks

	GetAllFriends(userID uuid.UUID) ([]UserInfo, error) //status true
	GetAllIncomingRequests(userID uuid.UUID) ([]UserInfo, error)
	GetAllSentRequests(userID uuid.UUID) ([]UserInfo, error)
	SendFriendRequest(userID, defID uuid.UUID) error
	ConfirmFriendRequest(userID, defID uuid.UUID) error
	DeleteIncomingRequest(userID, defID uuid.UUID) error
	DeleteSentRequest(userID, defID uuid.UUID) error

	GetAllBlockedUsers(userID uuid.UUID) ([]UserInfo, error)
	CreateBlock(userID, defID uuid.UUID) error
	DeleteBlock(userID, defID uuid.UUID) error
}

type FileStorages interface {
	GetAllStorageFiles(userID uuid.UUID) ([]File, error)

	GetAllImageFiles(userID uuid.UUID) ([]File, error)
	GetAllVideoFiles(userID uuid.UUID) ([]File, error)
	GetAllOtherFiles(userID uuid.UUID) ([]File, error)

	GetStorageFile(userID, fileID uuid.UUID) (File, error)

	UploadImageFile(userID, fileID uuid.UUID) error
	UploadVideoFile(userID, fileID uuid.UUID) error
	UploadOtherFile(userID, fileID uuid.UUID) error

	DeleteStorageFile(userID, fileID uuid.UUID) error
}

type ChatLists interface {
	GetAllChats(userID uuid.UUID) ([]ChatListItem, error) //without hided //By membership
	GetAllArchivedChats(userID uuid.UUID) ([]ChatListItem, error)

	ArchiveChat(userID, chatID uuid.UUID) error
	RemoveChatFromArchive(userID, chatID uuid.UUID) error

	GetChat(userID, chatID uuid.UUID) (ChatItem, error)
	CreateChat(userID, defID uuid.UUID) error //non-admin

	AddChatMembers(userID uuid.UUID, defIDs []uuid.UUID) error

	RemoveChatMember(userID, defIDs uuid.UUID) error
	DeleteChat(userID, chatID uuid.UUID) error

	HideChat(userID, chatID uuid.UUID) error
	CleanChatHistory(userID, chatID uuid.UUID)

	CreateGroupChat(userID uuid.UUID, chat Chat, file File) error //with admin

	CreateGroupChatAvatar(userID uuid.UUID, chat Chat, file File) error //By files
	UpdateGroupChatAvatar(userID, chatID uuid.UUID, file File) error

	UpdateGroupChatTitle(userID, chatID uuid.UUID, title string) error

	SendMessage(userID uuid.UUID, message Message) error
	ForwardMessage(userID uuid.UUID, message Message) error
	DeleteMessage(userID, messageID uuid.UUID) error

	LeaveChatGroup(userID, chatID uuid.UUID) error //non-group chat can be hided only
}

type Files interface {
	GetObjectLink(fileName string) string
	UploadObject(ctx context.Context, file File) (string, error)
	RemoveObject(ctx context.Context, object string) error
	RemoveFile(filename string)
}

type Services struct {
	Auths Auths
}

type Deps struct {
	Repos           *repository.Repositories
	Hasher          hash.PasswordHasher
	TokenManager    auth.TokenManager
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	StorageProvider storage.Provider
}

func NewServices(deps Deps) *Services {
	authsService := NewAuthsService(deps.Repos.Users, deps.Repos.Sessions, deps.Hasher, deps.TokenManager, deps.AccessTokenTTL, deps.RefreshTokenTTL)

	return &Services{
		Auths: authsService,
	}
}
