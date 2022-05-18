package service

import (
	"context"
	"errors"

	"github.com/artomsopun/clendry/clendry-api/internal/repository"
	"github.com/artomsopun/clendry/clendry-api/pkg/files"
	"github.com/artomsopun/clendry/clendry-api/pkg/hash"
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
)

type ProfilesService struct {
	repoUsers repository.Users
	repoFiles repository.Files
	files     files.Files
	hasher    hash.PasswordHasher
}

func NewProfilesService(repoUsers repository.Users, repoFiles repository.Files, files files.Files, hasher hash.PasswordHasher) *ProfilesService {
	return &ProfilesService{
		repoUsers: repoUsers,
		repoFiles: repoFiles,
		files:     files,
		hasher:    hasher,
	}
}

func (s *ProfilesService) GetProfile(userID types.BinaryUUID) (UserInfo, error) {
	user, err := s.repoUsers.GetById(userID)
	if err != nil {
		return UserInfo{}, err
	}
	file, err := s.repoFiles.GetAvatarByUserID(userID)
	if err != nil {
		return UserInfo{}, err
	}
	return UserInfo{
		ID:     user.ID,
		Nick:   user.Nick,
		Email:  user.Email,
		Avatar: s.files.GetObjectLink(file.Title),
	}, nil
}

func (s *ProfilesService) ChangePassword(confirm PasswordConfirm) error {
	user, err := s.repoUsers.GetById(confirm.UserID)
	if err != nil {
		return err
	}
	oldHash, err := s.hasher.Hash(confirm.OldPassword)
	if err != nil {
		return err
	}
	if user.Password != oldHash {
		return errors.New("wrong password")
	}
	if confirm.Passwords.Password != confirm.Passwords.Confirm {
		return errors.New("passwords mismatch")
	}
	passwordHash, err := s.hasher.Hash(confirm.Passwords.Confirm)
	if err != nil {
		return err
	}
	err = s.repoUsers.ChangePassword(confirm.UserID, passwordHash)
	return err
}

func (s *ProfilesService) DeleteProfile(userID types.BinaryUUID) error {
	err := s.repoUsers.Delete(userID)
	return err
}

func (s *ProfilesService) UploadAvatar(ctx context.Context, file File) error {
	// fileName, err := s.files.UploadObject(ctx, users, file)
	// if err != nil {
	// 	return err
	// }

	// err = s.repoFiles.CreateAvatarByUserID(domain.File{
	// 	Title:       fileName,
	// 	Size:        file.Size,
	// 	Current:     true,
	// 	ContentType: file.ContentType,
	// 	Type:        file.Type,
	// 	CreatedAt:   time.Now(),
	// 	UserID:      file.ForeignID,
	// })
	// return err
	return nil
}

func (s *ProfilesService) ChangeAvatarByFileID(userID, fileID types.BinaryUUID) error {
	err := s.repoFiles.ChangeAvatarByUserID(userID, fileID)
	return err
}
