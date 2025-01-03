package service

import (
	"ecommerce/models"
	"ecommerce/repository"
)

type ProductService struct {
	repository repository.Product
}

type Product interface {
	CreateProduct(user models.CreateProductDTO) error
	UpdateProduct(user models.UpdateProductDTO) error
	DeleteProduct(id string) error
	GetProducts(name, description string) ([]models.Product, error)
}

func NewProductService(repository repository.Product) *ProductService {
	return &ProductService{
		repository,
	}
}

func (p *ProductService) CreateProduct(product models.CreateProductDTO) error {
	return p.repository.InsertProduct(product)
}

func (p *ProductService) UpdateProduct(product models.UpdateProductDTO) error {
	return p.repository.UpdateProduct(product)
}

func (p *ProductService) DeleteProduct(id string) error {
	return p.repository.DeleteProduct(id)
}

func (p *ProductService) GetProducts(name, description string) ([]models.Product, error) {
	return p.repository.GetProducts(name, description)
}
