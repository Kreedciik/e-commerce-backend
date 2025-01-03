package service

import (
	"ecommerce/repository"
)

type Service struct {
	User
	Product
	Cart
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		User:    NewUserService(repository),
		Product: NewProductService(repository),
		Cart:    NewCartService(repository, repository.Product),
	}
}
