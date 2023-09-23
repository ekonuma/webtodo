package usecase

import (
	"os"
	"time"

	"github.com/ekonuma/webtodo/model"
	"github.com/ekonuma/webtodo/repository"
	"github.com/ekonuma/webtodo/validator"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	LogIn(user model.User) (string, error)
}

type userUsecase struct {
	repository    repository.IUserRepository
	userValidator validator.IUserValidator
}

func NewUserUserCase(repository repository.IUserRepository, userValidator validator.IUserValidator) IUserUsecase {
	return &userUsecase{repository, userValidator}
}

func (userUsecase *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	if err := userUsecase.userValidator.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}
	
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{Email: user.Email, Password: string(hash)}
	if err := userUsecase.repository.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (userUsecase *userUsecase) LogIn(user model.User) (string, error) {
	if err := userUsecase.userValidator.UserValidate(user); err != nil {
		return "", err
	}

	storedUser := model.User{}
	if err := userUsecase.repository.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
