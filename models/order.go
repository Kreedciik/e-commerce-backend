package models

import "time"

type Order struct {
	Id         string    `json:"id"`
	UserId     string    `json:"userId"`
	TotalPrice float64   `json:"totalPrice"`
	CreatedAt  time.Time `json:"-"`
}

type OrderItem struct {
	Id        string    `json:"id"`
	OrderId   string    `json:"orderId"`
	ProductId string    `json:"productId"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"-"`
}

type OrderCreateDTO struct {
	UserId     string  `json:"userId" binding:"required"`
	TotalPrice float64 `json:"totalPrice" binding:"required"`
}

type OrderItemCreateDTO struct {
	Id        string
	OrderId   string
	ProductId string
	Quantity  int
	Price     float64
}

type OrderUpdateDTO = Order
