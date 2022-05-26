package service

import (
	"context"
	"time"

	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"github.com/artomsopun/clendry/clendry-api/internal/repository"
	"github.com/artomsopun/clendry/clendry-api/pkg/auth"
	"github.com/artomsopun/clendry/clendry-api/pkg/files"
	"github.com/artomsopun/clendry/clendry-api/pkg/hash"
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
)

type UserInfo struct {
	ID       types.BinaryUUID
	Nick     string
	Email    string
	Avatar   string
	BlockTo  bool
	BlockBy  bool
	Sent     bool
	Incoming bool
	Friend   bool
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
	ID types.BinaryUUID

	Title       string
	Url         string
	Size        int64
	ContentType string
	Type        domain.FileType
	IsFavourite bool
	IsTrash     bool
	CreatedAt   time.Time

	MemberID types.BinaryUUID
	UserID   types.BinaryUUID
}

const (
	USERS    = "users"
	CHATS    = "chats"
	MESSAGES = "messages"
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

type Folder struct {
	ID types.BinaryUUID

	Title     string
	CreatedAt time.Time

	UserID types.BinaryUUID
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
	GetAllUsers(userID types.BinaryUUID) ([]UserInfo, error)
	GetUserByID(userID, defID types.BinaryUUID) (UserInfo, error)

	GetAllFriends(userID types.BinaryUUID) ([]UserInfo, error) //status true
	GetAllIncomingRequests(userID types.BinaryUUID) ([]UserInfo, error)
	GetAllSentRequests(userID types.BinaryUUID) ([]UserInfo, error)

	SendFriendRequest(userID, defID types.BinaryUUID) error
	ConfirmFriendRequest(userID, defID types.BinaryUUID) error

	DeleteRequest(userID, defID types.BinaryUUID) error

	GetAllBlockedUsers(userID types.BinaryUUID) ([]UserInfo, error)

	CreateBlock(userID, defID types.BinaryUUID) error
	DeleteBlock(userID, defID types.BinaryUUID) error
}

type FileStorages interface {
	GetAllFiles(userID types.BinaryUUID) ([]File, error)
	GetAllFilesByType(userID types.BinaryUUID, fileType domain.FileType) ([]File, error)
	GetFile(userID, fileID types.BinaryUUID) (File, error)
	GetContentType(ctype string) domain.FileType
	UploadFile(ctx context.Context, userID string, file File) (types.BinaryUUID, error)
	ChangeFileTitle(userID, fileID types.BinaryUUID, title string) error

	GetAllFilesByFolderID(userID, folderID types.BinaryUUID) ([]File, error)
	AddFileToFolder(userID, folderID, fileID types.BinaryUUID) error
	DeleteFileFromFolder(memberID types.BinaryUUID) error

	DeleteFile(userID, fileID types.BinaryUUID) error
	RemoveFile(filename string)

	GetAllFavouriteByUserID(userID types.BinaryUUID) ([]File, error)
	AddToFavourite(userID, fileID types.BinaryUUID) error
	DeleteFromFavourite(userID, fileID types.BinaryUUID) error

	GetAllTrashByUserID(userID types.BinaryUUID) ([]File, error)
	AddToTrash(userID, fileID types.BinaryUUID) error
	DeleteFromTrash(userID, fileID types.BinaryUUID) error
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

type Folders interface {
	GetAllFoldersByUserID(userID types.BinaryUUID) ([]Folder, error)
	GetFolderByUserID(userID, folderID types.BinaryUUID) (Folder, error)
	ChangeFolderTitleUserID(userID, folderID types.BinaryUUID, title string) error
	CreateFolder(folder Folder) error
	DeleteFolderByID(userID, folderID types.BinaryUUID) error
}

type Services struct {
	Auths    Auths
	Profiles Profiles
	Users    Users
	Storages FileStorages
	Folders  Folders
}

type Deps struct {
	Repos           *repository.Repositories
	Hasher          hash.PasswordHasher
	FilesManager    files.Files
	TokenManager    auth.TokenManager
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func NewServices(deps Deps) *Services {
	authsService := NewAuthsService(deps.Repos.Users, deps.Repos.Sessions, deps.Hasher, deps.TokenManager, deps.AccessTokenTTL, deps.RefreshTokenTTL)
	profilesService := NewProfilesService(deps.Repos.Users, deps.Repos.Files, deps.FilesManager, deps.Hasher)
	//usersService := NewUsersService(deps.Repos.Users, deps.Repos.Files, deps.FilesManager)//, deps.Repos.BlockRequests, deps.Repos.FriendRequests,)
	fileStorage := NewFilesService(deps.Repos.Files, deps.Repos.FolderFiles, deps.FilesManager)
	foldersService := NewFoldersService(deps.Repos.Folders)

	return &Services{
		Auths:    authsService,
		Profiles: profilesService,
		//Users:    usersService,
		Storages: fileStorage,
		Folders:  foldersService,
	}
}
