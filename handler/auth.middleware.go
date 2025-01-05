package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CheckAuthMiddleware(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		NewErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}
	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		NewErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}

	tokenString := authToken[1]
	claims, err := h.services.User.ParseToken(tokenString)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Set("currentUser", claims)
	ctx.Next()
}
