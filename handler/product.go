package handler

import (
	"ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateProduct(ctx *gin.Context) {
	var product models.CreateProductDTO

	if err := ctx.ShouldBindBodyWithJSON(&product); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.Product.CreateProduct(product); err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	NewSuccessResponse(ctx, "product successfully created")
}

func (h *Handler) UpdateProduct(ctx *gin.Context) {
	var product models.UpdateProductDTO
	if err := ctx.ShouldBindBodyWithJSON(&product); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.Product.UpdateProduct(product); err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	NewSuccessResponse(ctx, "product successfully updated")
}
func (h *Handler) GetProducts(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "")
	description := ctx.DefaultQuery("description", "")

	products, err := h.services.Product.GetProducts(name, description)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	NewSuccessResponse(ctx, products)
}
func (h *Handler) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.services.Product.DeleteProduct(id); err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(ctx, "successfully deleted")
}
