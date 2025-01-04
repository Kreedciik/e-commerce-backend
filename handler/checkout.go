package handler

import (
	"ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Checkout(ctx *gin.Context) {
	var checkoutDetails models.CreateCheckoutDTO
	if err := ctx.ShouldBindBodyWithJSON(&checkoutDetails); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if checkoutDetails.UserId == "" {
		NewErrorResponse(ctx, http.StatusBadRequest, "user ID is empty")
		return
	}

	if err := h.services.Order.CreateOrder(checkoutDetails); err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(ctx, "order successfully created")

}
