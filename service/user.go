package service

import (
	"ecommerce/models"
	"ecommerce/repository"
)

type UserService struct {
	repository repository.User
}

// All methods of user service should be implemented here
type User interface {
	CreateUser(user models.CreateUserDTO) error
	UpdateUser(user models.UpdateUserDTO) error
	DeleteUser(id string) error
	GetUserByID(id string) (models.User, error)
}

func NewUserService(repository repository.User) *UserService {
	return &UserService{
		repository,
	}
}

func (u *UserService) CreateUser(user models.CreateUserDTO) error {
	return u.repository.InsertUser(user)
}

func (u *UserService) UpdateUser(user models.UpdateUserDTO) error {
	return u.repository.UpdateUser(user)
}

func (u *UserService) DeleteUser(id string) error {
	return u.repository.DeleteUser(id)
}

func (u *UserService) GetUserByID(id string) (models.User, error) {
	return u.repository.FindUserByID(id)
}
