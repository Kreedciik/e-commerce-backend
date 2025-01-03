package service

import (
	"ecommerce/models"
	"ecommerce/repository"
	"errors"
)

type CartService struct {
	repository  repository.Cart
	productRepo repository.Product
}

type Cart interface {
	PutToCart(item models.PutToCartDTO) error
	RemoveFromCart(userId, productId string) error
	GetProductsFromCart(userId string) ([]models.Product, error)
}

func NewCartService(repository repository.Cart, productRepo repository.Product) *CartService {
	return &CartService{
		repository,
		productRepo,
	}
}

func (c *CartService) PutToCart(item models.PutToCartDTO) error {
	product, err := c.productRepo.FindProductById(item.ProductId)
	if err != nil {
		return err
	}
	if product.Stock < item.Quantity {
		return errors.New("the quantity exceeds")
	}
	return c.repository.InsertToCart(item)
}

func (c *CartService) RemoveFromCart(userId, productId string) error {
	return c.repository.RemoveFromCart(userId, productId)
}

func (c *CartService) GetProductsFromCart(userId string) ([]models.Product, error) {
	return c.repository.FindAllProductsFromCart(userId)
}
