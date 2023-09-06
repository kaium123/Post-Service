package service

import (
	"auth/models"
	"auth/repository"
)

type UserServiceInterface interface {
	Register(user models.User) error
	SignIn(signInInfo models.SignInData) error
}

type UserService struct {
	repository repository.UserRepositoryInterface
}

func NewUserService(repository repository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{repository: repository}
}

func (service *UserService) Register(user models.User) error {
	
    return service.repository.Register(user)
}

func (service *UserService) SignIn(signInInfo models.SignInData) error {
	
    return service.repository.Register(signInInfo)
}