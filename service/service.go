package service

import (
	"ecommerce/repository"
)

type Service struct {
	User
	Product
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		User:    NewUserService(repository),
		Product: NewProductService(repository),
	}
}
