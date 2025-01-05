package service

import (
	"ecommerce/models"
	"ecommerce/repository"

	"github.com/google/uuid"
)

type OrderService struct {
	repository  repository.Order
	productRepo repository.Product
	cartRepo    repository.Cart
}

type Order interface {
	CreateOrder(checkoutDetail models.CreateCheckoutDTO) error
}

func NewOrderService(
	repository repository.Order,
	productRepo repository.Product,
	cartRepo repository.Cart,
) *OrderService {
	return &OrderService{
		repository,
		productRepo,
		cartRepo,
	}
}

func (o *OrderService) CreateOrder(checkoutDetail models.CreateCheckoutDTO) error {
	var (
		totalPrice float64
		orderItems []models.OrderItemCreateDTO
		newOrderId = uuid.NewString()
	)
	userId := checkoutDetail.UserId
	products, err := o.cartRepo.FindAllProductsFromCart(userId)
	if err != nil {
		return err
	}
	for _, product := range products {
		totalPrice += product.Price * float64(product.Quantity)
		orderItems = append(orderItems, models.OrderItemCreateDTO{
			Id:        uuid.NewString(),
			OrderId:   newOrderId,
			ProductId: product.Id,
			Quantity:  product.Quantity,
			Price:     product.Price,
		})
	}

	newOrder := models.OrderCreateDTO{
		UserId:     userId,
		TotalPrice: totalPrice,
	}

	return o.repository.InsertOrder(newOrder, newOrderId, orderItems)
}
