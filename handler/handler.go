package handler

import (
	"ecommerce/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService *service.UserService
}

func NewHandler() *Handler {
	return &Handler{}
}

func Run() *gin.Engine {
	r := gin.Default()

	return r
}
