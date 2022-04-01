package repository

import (
	"errors"
	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"gorm.io/gorm"
)

type BlocksRepo struct {
	db *gorm.DB
}

func NewBlocksRepo(db *gorm.DB) *BlocksRepo {
	return &BlocksRepo{
		db: db,
	}
}

func (r *BlocksRepo) GetAllByUserID(userID uint) ([]domain.BlockRequest, error) {
	var blocks []domain.BlockRequest

	if err := r.db.Where("user_id = ?", userID).Find(&blocks).Error; err != nil {
		return blocks, err
	}

	return blocks, nil
}

func (r *BlocksRepo) Create(block domain.BlockRequest) error {
	err := r.db.Where("user_id = ? AND def_id = ?",
		block.UserID, block.DefID).First(&domain.BlockRequest{}).Error
	if err == nil {
		return errors.New("user already blocked")
	}

	r.db.Create(&block)
	return nil
}

func (r *BlocksRepo) Delete(userID, addresseeID uint) error {
	err := r.db.Where("user_id = ? AND def_id = ?", userID, addresseeID).Delete(&domain.BlockRequest{}).Error
	return err
}
