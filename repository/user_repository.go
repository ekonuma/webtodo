package repository

import (
	"github.com/ekonuma/webtodo/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository{
	return &userRepository{db}
}

func (userRepository *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := userRepository.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (userRepository *userRepository) CreateUser(user *model.User) error {
	if err := userRepository.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}