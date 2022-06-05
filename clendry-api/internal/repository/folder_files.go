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

func (r *FolderFilesRepo) GetByFolderFileID(folderID, fileID types.BinaryUUID) (domain.FolderFile, error) {
	var member domain.FolderFile
	if err := r.db.Where("folder_id = ? AND file_id = ?", folderID, fileID).First(&member).Error; err != nil {
		return domain.FolderFile{}, errors.New("files not found")
	}
	return member, nil
}

func (r *FolderFilesRepo) GetAllFilesByFolderID(folderID types.BinaryUUID) ([]domain.FolderFile, error) {
	var members []domain.FolderFile
	if err := r.db.Where("folder_id = ?", folderID).Find(&members).Error; err != nil {
		return []domain.FolderFile{}, errors.New("files not found")
	}
	return members, nil
}

func (r *FolderFilesRepo) GetAllFoldersByFileID(userID, fileID types.BinaryUUID) ([][]domain.Folder, error) {
	var allFolders [][]domain.Folder
	var foldersSelected []domain.Folder
	if err := r.db.Table("folders").Distinct().Select("folders.id, folders.title, folders.created_at, folders.user_id").
		Where("folders.user_id = ?", userID).
		Joins("left join folder_files on folder_files.folder_id = folders.id").
		Where("folder_files.file_id = ?", fileID).Scan(&foldersSelected).Error; err != nil {
		return [][]domain.Folder{}, errors.New("files not found")
	}
	var ids []types.BinaryUUID
	for _, folderSelected := range foldersSelected {
		ids = append(ids, folderSelected.ID)
	}
	allFolders = append(allFolders, foldersSelected)
	var foldersUnselected []domain.Folder
	if err := r.db.Where("user_id = ?", userID).Not(ids).Find(&foldersUnselected).Error; err != nil {
		return [][]domain.Folder{}, errors.New("files not found")
	}
	allFolders = append(allFolders, foldersUnselected)
	return allFolders, nil
}

func (r *FolderFilesRepo) GetAllNoFoldersByFileID(fileID types.BinaryUUID) ([]domain.FolderFile, error) {
	var members []domain.FolderFile
	if err := r.db.Where("file_id != ?", fileID).Find(&members).Error; err != nil {
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
