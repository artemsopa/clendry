package repository

import (
	"errors"
	"github.com/artomsopun/clendry/clendry-api/internal/domain"
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

func (r *FriendsRequestRepo) GetAllIncomingUnconfirmedByUserID(userID uint) ([]domain.FriendRequest, error) {
	var friends []domain.FriendRequest
	err := r.db.Where("addressee_id = ? AND status = ?", userID, false).Find(&friends).Error
	if err != nil {
		return []domain.FriendRequest{}, err
	}
	return friends, nil
}

func (r *FriendsRequestRepo) GetAllSentUnconfirmedByUserID(userID uint) ([]domain.FriendRequest, error) {
	var friends []domain.FriendRequest
	err := r.db.Where("user_id = ? AND status = ?", userID, false).Find(&friends).Error
	if err != nil {
		return []domain.FriendRequest{}, err
	}
	return friends, nil
}

func (r *FriendsRequestRepo) GetAllConfirmedByUserID(userID uint) ([]domain.FriendRequest, error) {
	var friends []domain.FriendRequest
	err := r.db.Where("user_id = ? AND status = ?", userID, true).Find(&friends).Error
	if err != nil {
		return []domain.FriendRequest{}, err
	}
	return friends, nil
}

func (r *FriendsRequestRepo) CreateUnconfirmed(request domain.FriendRequest) error {
	err := r.db.Where("user_id = ? AND def_id = ?",
		request.UserID, request.DefID).First(&domain.FriendRequest{}).Error
	if err == nil {
		return errors.New("request already created")
	}
	r.db.Create(&request)
	return nil
}

func (r *FriendsRequestRepo) UpdateConfirmation(request domain.FriendRequest) error {
	var friend domain.FriendRequest
	err := r.db.Where("user_id = ? AND def_id = ?",
		request.UserID, request.DefID).First(&friend).Error
	if err == nil {
		return errors.New("request doesn't exist")
	}
	if friend.Status == true {
		return errors.New("request had confirmed already")
	}
	err = r.db.Save(&request).Error
	return err
}

func (r *FriendsRequestRepo) DeleteRequest(userID, addresseeID uint) error {
	err := r.db.Where("user_id = ? AND def_id = ?", userID, addresseeID).Delete(&domain.FriendRequest{}).Error
	return err
}
