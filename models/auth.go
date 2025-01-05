package models

import "github.com/golang-jwt/jwt/v5"

type AuthDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserClaims struct {
	Role   string `json:"role"`
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}
