package repository

import (
	"errors"

	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"gorm.io/gorm"
)

type FolderFilesRepo struct {
	db *gorm.DB
}

func NewFolderFilesRepo(db *gorm.DB) *FolderFilesRepo {
	return &FolderFilesRepo{
		db: db,
	}
}

func (r *FolderFilesRepo) GetAllFilesByFolderID(folderID types.BinaryUUID) ([]domain.FolderFile, error) {
	var members []domain.FolderFile
	if err := r.db.Where("folder_id = ?", folderID).Find(&members).Error; err != nil {
		return []domain.FolderFile{}, errors.New("files not found")
	}
	return members, nil
}

func (r *FolderFilesRepo) CreateFolderFile(member domain.FolderFile) error {
	err := r.db.Create(&member).Error
	return err
}

func (r *FolderFilesRepo) DeleteByFolderFileID(memberID types.BinaryUUID) error {
	err := r.db.Where("id", memberID).Delete(&domain.FolderFile{}).Error
	return err
}
