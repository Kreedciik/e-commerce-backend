package models

type CreateCheckoutDTO struct {
	UserId string `json:"userId" binding:"required"`
}
