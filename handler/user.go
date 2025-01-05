package handler

import (
	"net/http"

	"ecommerce/config"
	"ecommerce/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(ctx *gin.Context) {
	user := models.CreateUserDTO{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
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
	user := models.UpdateUserDTO{}
	if err := ctx.BindJSON(&user); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.User.UpdateUser(user); err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(ctx, "User updated")
}
func (h *Handler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		NewErrorResponse(ctx, http.StatusBadRequest, "User ID is required")
		return
	}

	if err := h.services.User.DeleteUser(id); err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(ctx, "User deleted")
	// Implement delete user controller
}
func (h *Handler) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		NewErrorResponse(ctx, http.StatusBadRequest, "User ID is required")
		return
	}

	user, err := h.services.User.GetUserByID(id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if (user == models.User{}) {
		NewErrorResponse(ctx, http.StatusNotFound, "User not found")
		return
	}

	ctx.JSON(http.StatusOK, user)
	// Implement retrive user by id controller
}

func (h *Handler) LoginUser(ctx *gin.Context) {
	var authInput models.AuthDTO
	if err := ctx.ShouldBindJSON(&authInput); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	foundedUser, err := h.services.User.GetUserByEmail(authInput.Email)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "user not found")
		return
	}

	if err := h.services.User.ComparePasswords(foundedUser.Password, authInput.Password); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid password")
		return
	}

	accessToken, err := h.services.GenerateAccessToken(foundedUser, config.TOKEN_EXPIRE_DURATION, config.SECRET_KEY)

	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(ctx, gin.H{"accessToken": accessToken})

}
