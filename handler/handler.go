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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.CreateUser)
		auth.POST("/sign-in", h.LoginUser)
	}

	v1 := router.Group("/api/v1", h.CheckAuthMiddleware)
	{
		users := v1.Group("/users")
		{
			users.PUT("/update", h.UpdateUser)
			users.DELETE("/delete/:id", h.DeleteUser)
			users.GET("/:id", h.GetUserById)
		}

		products := v1.Group("/products", h.CheckAdminMiddleware)
		{
			products.POST("/create", h.CreateProduct)
			products.PUT("/update", h.UpdateProduct)
			products.GET("/", h.GetProducts)
			products.DELETE("/delete/:id", h.DeleteProduct)
		}

		carts := v1.Group("/carts")
		{
			carts.POST("/add", h.AddProductToCart)
			carts.DELETE("/:id", h.RemoveProductFromCart)
			carts.GET("/:id", h.GetProductsFromCart)
		}

		checkout := v1.Group("/checkout")
		{
			checkout.POST("", h.Checkout)
		}
	}
	return router
}
