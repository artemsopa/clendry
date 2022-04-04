package repository

import (
	"errors"
	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"gorm.io/gorm"
)

type FilesRepo struct {
	db *gorm.DB
}

func NewFilesRepo(db *gorm.DB) *FilesRepo {
	return &FilesRepo{
		db: db,
	}
}

func (r *FilesRepo) GetAllFilesByUserID(userID types.BinaryUUID) ([]domain.File, error) {
	var file []domain.File
	if err := r.db.Where("user_id = ? AND current = ?", userID, false).Find(&file).Error; err != nil {
		return []domain.File{}, errors.New("files not found")
	}
	return file, nil
}

func (r *FilesRepo) GetFileByUserID(userID, fileID types.BinaryUUID) (domain.File, error) {
	file := domain.File{}
	if err := r.db.Where("file_id = ? AND user_id = ?", fileID, userID).First(&file).Error; err != nil {
		return domain.File{}, errors.New("file not found")
	}
	return file, nil
}

func (r *FilesRepo) GetAvatarByUserID(userID types.BinaryUUID) (domain.File, error) {
	file := domain.File{}
	if err := r.db.Where("user_id = ? AND current = ?", userID, true).First(&file).Error; err != nil {
		return domain.File{}, errors.New("file not found")
	}
	return file, nil
}

func (r *FilesRepo) ChangeAvatarByUserID(userID, fileID types.BinaryUUID) error {
	err := r.db.Model(&domain.File{}).Where("user_id = ? ", userID).
		Update("current", false).Error
	if err != nil {
		return errors.New("you have not avatar")
	}
	err = r.db.Model(&domain.File{}).Where("id = ? AND user_id", fileID, userID).
		Update("current", true).Error
	if err != nil {
		return errors.New("you have not such file")
	}
	return nil
}

func (r *FilesRepo) CreateAvatarByUserID(file domain.File) error {
	err := r.db.Model(&domain.File{}).Where("user_id = ? ", file.UserID).
		Update("current", false).Error
	if err != nil {
		return errors.New("you have not avatar")
	}
	err = r.db.Create(&file).Error
	return err
}

func (r *FilesRepo) GetAllTypeFilesByUserID(userID types.BinaryUUID, filetype domain.FileType) ([]domain.File, error) {
	var files []domain.File
	if err := r.db.Where("user_id = ? AND type = ? AND current = ?", userID, filetype, false).Find(&files).Error; err != nil {
		return []domain.File{}, errors.New("files not found")
	}
	return files, nil
}

func (r *FilesRepo) Create(file domain.File) error {
	err := r.db.Create(&file).Error
	return err
}

func (r *FilesRepo) DeleteByID(userID, fileID types.BinaryUUID) error {
	if err := r.db.Where("file_id = ? AND user_id = ?", fileID, userID).First(&domain.File{}).Error; err != nil {
		return err
	}
	err := r.db.Delete(&domain.File{}, fileID).Error
	return err
}

/*
func (r *FilesRepo) GetAvatarByChatID(chatID types.BinaryUUID) (domain.File, error) {}

func (r *FilesRepo) ChangeChatAvatarByMemberID(memberID, chatID types.BinaryUUID) (domain.File, error) {
}

func (r *FilesRepo) CreateChatAvatarByMemberID(memberID, chatID types.BinaryUUID, file domain.File) error {}

func (r *FilesRepo) GetAllFilesByMessageID(messageID types.BinaryUUID) (domain.File, error) {}

func (r *FilesRepo) CreateFileByMessageID(chatID types.BinaryUUID) (domain.File, error) {}
*/
