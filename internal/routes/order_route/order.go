package order_route

import (
	"kcloudb1/internal/handlers/order_handler"
	"kcloudb1/internal/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoute(r *gin.RouterGroup) {
	orderRoute := r.Group("/order")
	{
		orderRoute.POST("/", middleware.Auth(), order_handler.CreateOrder)
		orderRoute.POST("/vip-member", middleware.Auth(), order_handler.CreateVIPMember)
		orderRoute.GET("/", middleware.Auth(), middleware.Paginate(), order_handler.GetOrderList)
		orderRoute.GET("/by-id", middleware.Auth(), order_handler.GetOrder)
	}
}
