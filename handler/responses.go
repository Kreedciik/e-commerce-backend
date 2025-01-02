package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{"error": message, "code": statusCode})
}

func NewSuccessResponse(ctx *gin.Context, message interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"message": message})
}
