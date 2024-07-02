package order_route

import (
	"kcloudb1/internal/handlers/order_handler"
	"kcloudb1/internal/middleware"

	"github.com/gin-gonic/gin"
)

func WithDrawRoute(r *gin.RouterGroup) {
	withdrawRoute := r.Group("/withdraw")
	{
		withdrawRoute.POST("/", middleware.Auth(), order_handler.CreateWithdraw)
		withdrawRoute.GET("/", middleware.Auth(), order_handler.GetWithDrawList)
		withdrawRoute.GET("/by-id", middleware.Auth(), order_handler.GetWithdraw)
	}
}
