package service

import (
	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"github.com/artomsopun/clendry/clendry-api/internal/repository"
	"github.com/artomsopun/clendry/clendry-api/pkg/files"
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
)

type UsersService struct {
	repoUsers   repository.Users
	repoBlocks  repository.BlockRequests
	repoFriends repository.FriendRequests
	repoFiles   repository.Files
	files       files.Files
}

func NewUsersService(repoUsers repository.Users, repoFiles repository.Files, repoBlocks repository.BlockRequests,
	repoFriends repository.FriendRequests, files files.Files) *UsersService {
	return &UsersService{
		repoUsers:   repoUsers,
		repoBlocks:  repoBlocks,
		repoFriends: repoFriends,
		repoFiles:   repoFiles,
		files:       files,
	}
}

func (s *UsersService) GetAllUsers(userID types.BinaryUUID) ([]UserInfo, error) {
	usersRepo, err := s.repoUsers.GetAllWithoutBlocks(userID)
	if err != nil {
		return nil, err
	}
	var users []UserInfo
	for _, user := range usersRepo {
		file, err := s.repoFiles.GetAvatarByUserID(user.ID)
		if err != nil {
			return nil, err
		}
		isBlockBy := s.repoBlocks.IsUserInBlock(userID, user.ID)
		isBlockTo := s.repoBlocks.IsDefInBlock(userID, user.ID)
		isSent := s.repoFriends.IsSentReq(userID, user.ID)
		isIncoming := s.repoFriends.IsSentReq(userID, user.ID)
		isFriend := s.repoFriends.IsUserInFriend(userID, user.ID)
		users = append(users, UserInfo{
			ID:       user.ID,
			Nick:     user.Nick,
			Email:    user.Email,
			Avatar:   s.files.GetObjectLink(file.Title),
			BlockBy:  isBlockBy,
			BlockTo:  isBlockTo,
			Sent:     isSent,
			Incoming: isIncoming,
			Friend:   isFriend,
		})
	}
	return users, nil
}

func (s *UsersService) GetUserByID(userID, defID types.BinaryUUID) (UserInfo, error) {
	userRepo, err := s.repoUsers.GetById(defID)
	if err != nil {
		return UserInfo{}, err
	}
	file, err := s.repoFiles.GetAvatarByUserID(userRepo.ID)
	if err != nil {
		return UserInfo{}, err
	}
	isBlockBy := s.repoBlocks.IsUserInBlock(userID, userRepo.ID)
	isBlockTo := s.repoBlocks.IsDefInBlock(userID, userRepo.ID)
	isSent := s.repoFriends.IsSentReq(userID, userRepo.ID)
	isIncoming := s.repoFriends.IsSentReq(userID, userRepo.ID)
	isFriend := s.repoFriends.IsUserInFriend(userID, userRepo.ID)
	user := UserInfo{
		ID:       userRepo.ID,
		Nick:     userRepo.Nick,
		Email:    userRepo.Email,
		Avatar:   s.files.GetObjectLink(file.Title),
		BlockBy:  isBlockBy,
		BlockTo:  isBlockTo,
		Sent:     isSent,
		Incoming: isIncoming,
		Friend:   isFriend,
	}
	return user, nil
}

func (s *UsersService) GetAllFriends(userID types.BinaryUUID) ([]UserInfo, error) {
	usersRepo, err := s.repoUsers.GetAllFriends(userID)
	if err != nil {
		return nil, err
	}
	var users []UserInfo
	for _, user := range usersRepo {
		file, err := s.repoFiles.GetAvatarByUserID(user.ID)
		if err != nil {
			return nil, err
		}
		users = append(users, UserInfo{
			ID:     user.ID,
			Nick:   user.Nick,
			Email:  user.Email,
			Avatar: s.files.GetObjectLink(file.Title),
		})
	}
	return users, nil
}

func (s *UsersService) GetAllIncomingRequests(userID types.BinaryUUID) ([]UserInfo, error) {
	usersRepo, err := s.repoUsers.GetAllIncomingReqs(userID)
	if err != nil {
		return nil, err
	}
	var users []UserInfo
	for _, user := range usersRepo {
		file, err := s.repoFiles.GetAvatarByUserID(user.ID)
		if err != nil {
			return nil, err
		}
		users = append(users, UserInfo{
			ID:     user.ID,
			Nick:   user.Nick,
			Email:  user.Email,
			Avatar: s.files.GetObjectLink(file.Title),
		})
	}
	return users, nil
}

func (s *UsersService) GetAllSentRequests(userID types.BinaryUUID) ([]UserInfo, error) {
	usersRepo, err := s.repoUsers.GetAllSentReqs(userID)
	if err != nil {
		return nil, err
	}
	var users []UserInfo
	for _, user := range usersRepo {
		file, err := s.repoFiles.GetAvatarByUserID(user.ID)
		if err != nil {
			return nil, err
		}
		users = append(users, UserInfo{
			ID:     user.ID,
			Nick:   user.Nick,
			Email:  user.Email,
			Avatar: s.files.GetObjectLink(file.Title),
		})
	}
	return users, nil
}

func (s *UsersService) SendFriendRequest(userID, defID types.BinaryUUID) error {
	err := s.repoFriends.CreateUnconfirmed(domain.FriendRequest{
		Status: false,
		UserID: userID,
		DefID:  defID,
	})
	return err
}

func (s *UsersService) ConfirmFriendRequest(userID, defID types.BinaryUUID) error {
	err := s.repoFriends.UpdateConfirmation(domain.FriendRequest{
		Status: true,
		UserID: defID,
		DefID:  userID,
	})
	return err
}

func (s *UsersService) DeleteIncomingRequest(userID, defID types.BinaryUUID) error {
	err := s.repoFriends.DeleteReq(defID, userID, false)
	return err
}

func (s *UsersService) DeleteSentRequest(userID, defID types.BinaryUUID) error {
	err := s.repoFriends.DeleteReq(userID, defID, false)
	return err
}

func (s *UsersService) DeleteConfirmFriendRequest(userID, defID types.BinaryUUID) error {
	err := s.repoFriends.DeleteReq(userID, defID, true)
	return err
}

func (s *UsersService) GetAllBlockedUsers(userID types.BinaryUUID) ([]UserInfo, error) {
	usersRepo, err := s.repoUsers.GetAllBlockedUsers(userID)
	if err != nil {
		return nil, err
	}
	var users []UserInfo
	for _, user := range usersRepo {
		file, err := s.repoFiles.GetAvatarByUserID(user.ID)
		if err != nil {
			return nil, err
		}
		users = append(users, UserInfo{
			ID:     user.ID,
			Nick:   user.Nick,
			Email:  user.Email,
			Avatar: s.files.GetObjectLink(file.Title),
		})
	}
	return users, nil
}

func (s *UsersService) CreateBlock(userID, defID types.BinaryUUID) error {
	err := s.repoBlocks.Create(domain.BlockRequest{
		UserID: userID,
		DefID:  defID,
	})
	return err
}

func (s *UsersService) DeleteBlock(userID, defID types.BinaryUUID) error {
	err := s.repoBlocks.Delete(userID, defID)
	return err
}
