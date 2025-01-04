package models

import "time"

type Cart struct {
	Id        string    `json:"id"`
	UserId    string    `json:"userId"`
	ProductId string    `json:"productId"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"-"`
}

type CartProduct struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Quantity    int     `json:"quantity"`
}

type PutToCartDTO struct {
	UserId    string `json:"userId" binding:"required"`
	ProductId string `json:"productId" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}
