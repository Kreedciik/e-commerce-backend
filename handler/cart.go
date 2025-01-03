package handler

import (
	"ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddProductToCart(ctx *gin.Context) {
	var item models.PutToCartDTO
	if err := ctx.ShouldBindBodyWithJSON(&item); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.Cart.PutToCart(item); err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	NewSuccessResponse(ctx, "product successfully added")
}
func (h *Handler) RemoveProductFromCart(ctx *gin.Context) {
	userId := ctx.DefaultQuery("userId", "")
	productId := ctx.DefaultQuery("productId", "")

	if userId == "" || productId == "" {
		NewErrorResponse(ctx, http.StatusBadRequest, "provide valid userId or productId")
		return
	}
	if err := h.services.Cart.RemoveFromCart(userId, productId); err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(ctx, "product successfully deleted")
}
func (h *Handler) GetProductsFromCart(ctx *gin.Context) {
	id := ctx.Param("id")
	products, err := h.services.Cart.GetProductsFromCart(id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	NewSuccessResponse(ctx, products)
}
