package handler

import (
	"ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CheckAdminMiddleware(ctx *gin.Context) {
	user, _ := ctx.Get("currentUser")
	userClaims := user.(*models.UserClaims)
	if userClaims.Role != "admin" {
		NewErrorResponse(ctx, http.StatusUnauthorized, "access denied")
		return
	}
	ctx.Next()
}
