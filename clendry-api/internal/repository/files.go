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

func (r *FilesRepo) GetFilesKBSum(userID types.BinaryUUID) int {
	var sum int
	if err := r.db.Table("files").Select("sum(size)").Where("user_id = ?", userID).Row().Scan(&sum); err != nil {
		return 0
	}
	return sum
}

func (r *FilesRepo) GetAllFilesByUserID(userID types.BinaryUUID) ([]domain.File, error) {
	var file []domain.File
	if err := r.db.Where("user_id = ? AND is_trash = ?", userID, false).Find(&file).Error; err != nil {
		return []domain.File{}, errors.New("files not found")
	}
	return file, nil
}

func (r *FilesRepo) GetFileByUserID(userID, fileID types.BinaryUUID) (domain.File, error) {
	file := domain.File{}
	if err := r.db.Where("id = ? AND user_id = ?", fileID, userID).First(&file).Error; err != nil {
		return domain.File{}, errors.New("file not found")
	}
	return file, nil
}

func (r *FilesRepo) GetFileByFolder(userID, fileID types.BinaryUUID) (domain.File, error) {
	file := domain.File{}
	if err := r.db.Where("id = ? AND user_id = ? AND is_trash = ?", fileID, userID, false).First(&file).Error; err != nil {
		return domain.File{}, errors.New("file not found")
	}
	return file, nil
}

func (r *FilesRepo) GetAvatarByUserID(userID types.BinaryUUID) (domain.File, error) {
	file := domain.File{}
	// if err := r.db.Where("user_id = ? AND current = ?", userID, true).First(&file).Error; err != nil {
	// 	return domain.File{}, errors.New("file not found")
	// }
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
	if err := r.db.Where("user_id = ? AND type = ? AND is_trash = ?", userID, filetype, false).Find(&files).Error; err != nil {
		return []domain.File{}, errors.New("files not found")
	}
	return files, nil
}

func (r *FilesRepo) Create(file domain.File) (types.BinaryUUID, error) {
	err := r.db.Create(&file).Error
	if err != nil {
		return types.BinaryUUID{}, err
	}
	return file.ID, nil
}

func (r *FilesRepo) UpdateUploads(userID types.BinaryUUID) error {
	user := domain.User{}
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return errors.New("user not found")
	}
	err := r.db.Model(&domain.User{}).Where("id = ?", userID).Update("uploads", user.Uploads+1).Error
	return err
}

func (r *FilesRepo) DeleteByID(userID, fileID types.BinaryUUID) error {
	err := r.db.Where("id = ? AND user_id = ?", fileID, userID).Delete(&domain.File{}).Error
	return err
}

func (r *FilesRepo) GetAllFilesByUserFolderID(userID, folderID types.BinaryUUID) ([]domain.File, error) {
	var files []domain.File
	if err := r.db.Where("user_id = ? AND folder_id = ? AND is_trash = ?", userID, folderID, false).Find(&files).Error; err != nil {
		return []domain.File{}, errors.New("files not found")
	}
	return files, nil
}

func (r *FilesRepo) AddToFolder(userID, folderID, fileID types.BinaryUUID) error {
	if err := r.db.Where("id = ? AND user_id = ?", fileID, userID).First(&domain.File{}).Error; err != nil {
		return err
	}
	err := r.db.Model(&domain.File{}).Where("id = ? AND user_id = ?", fileID, userID).
		Update("folder_id", folderID).Error
	return err
}

func (r *FilesRepo) DeleteFromFolder(userID, fileID types.BinaryUUID) error {
	if err := r.db.Where("id = ? AND user_id = ?", fileID, userID).First(&domain.File{}).Error; err != nil {
		return err
	}
	err := r.db.Model(&domain.File{}).Where("id = ? AND user_id = ?", fileID, userID).
		Update("folder_id", nil).Error
	return err
}

func (r *FilesRepo) GetAllFavouriteByUserID(userID types.BinaryUUID) ([]domain.File, error) {
	var file []domain.File
	if err := r.db.Where("user_id = ? AND is_favourite = ? AND is_trash = ?", userID, true, false).Find(&file).Error; err != nil {
		return []domain.File{}, errors.New("files not found")
	}
	return file, nil
}

func (r *FilesRepo) AddToFavourite(userID, fileID types.BinaryUUID) error {
	if err := r.db.Where("id = ? AND user_id = ?", fileID, userID).First(&domain.File{}).Error; err != nil {
		return err
	}
	err := r.db.Model(&domain.File{}).Where("id = ? AND user_id = ?", fileID, userID).
		Update("is_favourite", true).Error
	return err
}

func (r *FilesRepo) DeleteFromFavourite(userID, fileID types.BinaryUUID) error {
	if err := r.db.Where("id = ? AND user_id = ?", fileID, userID).First(&domain.File{}).Error; err != nil {
		return err
	}
	err := r.db.Model(&domain.File{}).Where("id = ? AND user_id = ?", fileID, userID).
		Update("is_favourite", false).Error
	return err
}

func (r *FilesRepo) GetAllTrashByUserID(userID types.BinaryUUID) ([]domain.File, error) {
	var file []domain.File
	if err := r.db.Where("user_id = ? AND is_trash = ?", userID, true).Find(&file).Error; err != nil {
		return []domain.File{}, errors.New("files not found")
	}
	return file, nil
}
func (r *FilesRepo) AddToTrash(userID, fileID types.BinaryUUID) error {
	if err := r.db.Where("id = ? AND user_id = ?", fileID, userID).First(&domain.File{}).Error; err != nil {
		return err
	}
	err := r.db.Model(&domain.File{}).Where("id = ? AND user_id = ?", fileID, userID).
		Update("is_trash", true).Error
	return err
}

func (r *FilesRepo) DeleteFromTrash(userID, fileID types.BinaryUUID) error {
	if err := r.db.Where("id = ? AND user_id = ?", fileID, userID).First(&domain.File{}).Error; err != nil {
		return err
	}
	err := r.db.Model(&domain.File{}).Where("id = ? AND user_id = ?", fileID, userID).
		Update("is_trash", false).Error
	return err
}

func (r *FilesRepo) ChangeFileTitle(userID, fileID types.BinaryUUID, title string) error {
	err := r.db.Where("title = ?", title).First(&domain.File{}).Error
	if err == nil {
		return errors.New("title already exists")
	}
	err = r.db.Model(&domain.File{}).Where("id = ? AND user_id = ?", fileID, userID).
		Updates(&domain.File{Title: title}).Error
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
