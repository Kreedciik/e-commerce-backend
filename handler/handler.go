package handler

import (
	"ecommerce/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/sign-up", h.CreateUser)
			users.PUT("/update", h.UpdateUser)
			users.DELETE("/delete/:id", h.DeleteUser)
			users.GET("/:id", h.GetUserById)
		}
	}
	return router
}
