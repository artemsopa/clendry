package repository

import (
	"errors"

	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"gorm.io/gorm"
)

type FoldersRepo struct {
	db *gorm.DB
}

func NewFoldersRepo(db *gorm.DB) *FoldersRepo {
	return &FoldersRepo{
		db: db,
	}
}

func (r *FoldersRepo) GetAllFoldersByUserID(userID types.BinaryUUID) ([]domain.Folder, error) {
	var folders []domain.Folder
	if err := r.db.Where("user_id = ?", userID).Find(&folders).Error; err != nil {
		return []domain.Folder{}, errors.New("folders not found")
	}
	return folders, nil
}

func (r *FoldersRepo) GetFolderByUserID(userID, folderID types.BinaryUUID) (domain.Folder, error) {
	folder := domain.Folder{}
	if err := r.db.Where("id = ? AND user_id = ?", folderID, userID).First(&folder).Error; err != nil {
		return domain.Folder{}, errors.New("folder not found")
	}
	return folder, nil
}

func (r *FoldersRepo) ChangeFolderTitleUserID(userID, folderID types.BinaryUUID, title string) error {
	err := r.db.Model(&domain.Folder{}).Where("id = ? AND user_id = ? ", folderID, userID).
		Update("title", title).Error
	return err
}

func (r *FoldersRepo) Create(folder domain.Folder) error {
	err := r.db.Create(&folder).Error
	return err
}

func (r *FoldersRepo) DeleteByID(userID, folderID types.BinaryUUID) error {
	err := r.db.Where("id = ? AND user_id = ?", folderID, userID).Delete(&domain.Folder{}).Error
	return err
}
