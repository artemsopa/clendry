package repository

import (
	"errors"
	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"gorm.io/gorm"
)

type BlocksRequestRepo struct {
	db *gorm.DB
}

func NewBlocksRepo(db *gorm.DB) *BlocksRequestRepo {
	return &BlocksRequestRepo{
		db: db,
	}
}

func (r *BlocksRequestRepo) Create(block domain.BlockRequest) error {
	err := r.db.Where("user_id = ? AND def_id = ?",
		block.UserID, block.DefID).First(&domain.BlockRequest{}).Error
	if err == nil {
		return errors.New("user already blocked")
	}

	r.db.Create(&block)
	return nil
}

func (r *BlocksRequestRepo) Delete(userID, addresseeID types.BinaryUUID) error {
	err := r.db.Where("user_id = ? AND def_id = ?", userID, addresseeID).Delete(&domain.BlockRequest{}).Error
	return err
}

func (r *BlocksRequestRepo) IsDefInBlock(userID, defID types.BinaryUUID) bool {
	if err := r.db.Where("user_id = ? AND def_id = ?", userID, defID).
		First(&domain.BlockRequest{}).Error; err != nil {
		return false
	}
	return true
}

func (r *BlocksRequestRepo) IsUserInBlock(userID, defID types.BinaryUUID) bool {
	if err := r.db.Where("user_id = ? AND def_id = ?", defID, userID).
		First(&domain.BlockRequest{}).Error; err != nil {
		return false
	}
	return true
}
