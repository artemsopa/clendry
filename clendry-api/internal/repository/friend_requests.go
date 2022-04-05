package repository

import (
	"errors"
	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"gorm.io/gorm"
)

type FriendsRequestRepo struct {
	db *gorm.DB
}

func NewFriendsRepo(db *gorm.DB) *FriendsRequestRepo {
	return &FriendsRequestRepo{
		db: db,
	}
}

func (r *FriendsRequestRepo) CreateUnconfirmed(request domain.FriendRequest) error {
	err := r.db.Where("(user_id = ? AND def_id = ?) OR (user_id = ? AND def_id = ?)",
		request.UserID, request.DefID, request.DefID, request.UserID).
		First(&domain.FriendRequest{}).Error
	if err == nil {
		return errors.New("request already created")
	}
	r.db.Create(&request)
	return nil
}

func (r *FriendsRequestRepo) UpdateConfirmation(request domain.FriendRequest) error {
	var friend domain.FriendRequest
	err := r.db.Where("(user_id = ? AND def_id = ?) OR (user_id = ? AND def_id = ?)",
		request.UserID, request.DefID, request.DefID, request.UserID).
		First(&friend).Error
	if err != nil {
		return errors.New("request doesn't exist")
	}
	if friend.Status == true {
		return errors.New("request had confirmed already")
	}
	err = r.db.Save(&request).Error
	return err
}

func (r *FriendsRequestRepo) DeleteReq(userID, defID types.BinaryUUID, status bool) error {
	err := r.db.Where("(user_id = ? AND def_id = ?) OR (user_id = ? AND def_id = ?) AND status = ?",
		userID, defID, defID, userID, status).
		Delete(&domain.FriendRequest{}).Error
	return err
}

func (r *FriendsRequestRepo) IsUserInFriend(userID, defID types.BinaryUUID) bool {
	if err := r.db.Where(
		"((user_id = ? AND def_id = ?) OR (user_id = ? AND def_id = ?)) AND status = ?",
		userID, defID, defID, userID, true).
		First(&domain.FriendRequest{}).Error; err != nil {
		return false
	}
	return true
}

func (r *FriendsRequestRepo) IsSentReq(userID, defID types.BinaryUUID) bool {
	if err := r.db.Where(
		"user_id = ? AND def_id = ? AND status = ?",
		userID, defID, false).
		First(&domain.FriendRequest{}).Error; err != nil {
		return false
	}
	return true
}

func (r *FriendsRequestRepo) IsIncomingReq(userID, defID types.BinaryUUID) bool {
	if err := r.db.Where(
		"user_id = ? AND def_id = ? AND status = ?",
		defID, userID, false).
		First(&domain.FriendRequest{}).Error; err != nil {
		return false
	}
	return true
}
