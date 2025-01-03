package models

import "time"

type Cart struct {
	Id        string    `json:"id"`
	UserId    string    `json:"userId"`
	ProductId string    `json:"productId"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"-"`
}

type PutToCartDTO struct {
	UserId    string `json:"userId" binding:"required"`
	ProductId string `json:"productId" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}
