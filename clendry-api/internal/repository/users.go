package repository

import (
	"errors"
	"github.com/artomsopun/clendry/clendry-api/internal/domain"
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

func (r *UsersRepo) GetAll() ([]domain.User, error) {
	//TODO get all without blocks
	var users []domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

func (r *UsersRepo) GetById(userID uint) (domain.User, error) {
	//TODO if block return err
	user := domain.User{}
	if err := r.db.First(&user, userID).Error; err != nil {
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

func (r *UsersRepo) ChangePassword(userID uint, password string) error {
	err := r.db.Model(&domain.User{}).Where("id = ?", userID).Update("password", password).Error
	return err
}

func (r *UsersRepo) ChangeAvatar(userID uint, url string) error {
	err := r.db.Model(&domain.User{}).Where("id = ?", userID).Update("avatar", url).Error
	return err
}

func (r *UsersRepo) Delete(userID uint) error {
	err := r.db.Delete(&domain.User{}, userID).Error
	return err
}
