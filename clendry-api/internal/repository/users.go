package repository

import (
	"errors"

	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"gorm.io/gorm"
)

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) GetAll(userID types.BinaryUUID) ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Where("user_id != ?", userID).Find(&users).Error; err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

func (r *UsersRepo) GetAllWithoutBlocks(userID types.BinaryUUID) ([]domain.User, error) {
	var users []domain.User
	err := r.db.Model(&domain.User{}).Select("users.id, users.nick, users.email").
		Where("users.id != ?", userID).
		Joins("left join block_requests AS b on b.user_id = users.id AND b.def_id = users.id").
		Where("b.user_id != ? AND b.def_id != ?", userID, userID).
		Scan(&users).Error
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

func (r *UsersRepo) GetAllBlockedUsers(userID types.BinaryUUID) ([]domain.User, error) {
	var users []domain.User
	err := r.db.Model(&domain.User{}).Select("users.id, users.nick, users.email").
		//Where("users.id != ?", userID).
		Joins("left join block_requests AS b on b.user_id = users.id AND b.def_id = users.id").
		Where("b.user_id = ?", userID).
		Scan(&users).Error
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

func (r *UsersRepo) GetAllFriends(userID types.BinaryUUID) ([]domain.User, error) {
	var users []domain.User
	err := r.db.Model(&domain.User{}).Select("users.id, users.nick, users.email").
		//Where("users.id != ?", userID).
		Joins("left join friend_requests AS f on f.user_id = users.id AND f.def_id = users.id").
		Where("(f.user_id = ? OR f.def_id != ?) AND status = ?", userID, userID, true).
		Scan(&users).Error
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

func (r *UsersRepo) GetAllSentReqs(userID types.BinaryUUID) ([]domain.User, error) {
	var users []domain.User
	err := r.db.Model(&domain.User{}).Select("users.id, users.nick, users.email").
		//Where("users.id != ?", userID).
		Joins("left join friend_requests AS f on f.user_id = users.id AND f.def_id = users.id").
		Where("f.user_id = ? AND status = ?", userID, false).
		Scan(&users).Error
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

func (r *UsersRepo) GetAllIncomingReqs(userID types.BinaryUUID) ([]domain.User, error) {
	var users []domain.User
	err := r.db.Model(&domain.User{}).Select("users.id, users.nick, users.email").
		//Where("users.id != ?", userID).
		Joins("left join friend_requests AS f on f.user_id = users.id AND f.def_id = users.id").
		Where("f.def_id != ? AND status = ?", userID, false).
		Scan(&users).Error
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

func (r *UsersRepo) GetById(userID types.BinaryUUID) (domain.User, error) {
	user := domain.User{}
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (r *UsersRepo) GetByCredentials(login, password string) (domain.User, error) {
	user := domain.User{}
	if err := r.db.Where("(nick = ? OR email = ?) AND password = ?", login, login, password).First(&user).Error; err != nil {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (r *UsersRepo) Create(user domain.User) error {
	err := r.db.Where("email = ?", user.Email).First(&domain.User{}).Error
	if err == nil {
		return errors.New("user already exists")
	}
	err = r.db.Where("nickname = ?", user.Nick).First(&domain.User{}).Error
	if err == nil {
		return errors.New("user already exists")
	}
	r.db.Create(&user)
	return nil
}

func (r *UsersRepo) ChangePassword(userID types.BinaryUUID, password string) error {
	err := r.db.Model(&domain.User{}).Where("id = ?", userID).Update("password", password).Error
	return err
}

func (r *UsersRepo) Delete(userID types.BinaryUUID) error {
	err := r.db.Delete(&domain.User{}, userID).Error
	return err
}

func (r *UsersRepo) UpdateDownloads(userID types.BinaryUUID) error {
	user := domain.User{}
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return errors.New("user not found")
	}
	err := r.db.Model(&domain.User{}).Where("id = ?", userID).Update("downloads", user.Downloads+1).Error
	return err
}

func (r *UsersRepo) UpdateMemory(userID types.BinaryUUID, memory uint) error {
	err := r.db.Model(&domain.User{}).Where("id = ?", userID).Update("memory", memory).Error
	return err
}
