package validator

import "github.com/ekonuma/webtodo/model"

type IUserValidator interface {
	UserValidate(user model.User) error
}

type userValidator struct {}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}