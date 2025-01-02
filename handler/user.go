package handler

import (
	"net/http"

	"ecommerce/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(ctx *gin.Context) {
	user := models.CreateUserDTO{}
	if err := ctx.BindJSON(&user); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.User.CreateUser(user); err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(ctx, "User created")
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	// Implement update user controller
}
func (h *Handler) DeleteUser(ctx *gin.Context) {
	// Implement delete user controller
}
func (h *Handler) GetUserById(ctx *gin.Context) {
	// Implement retrive user by id controller
}
