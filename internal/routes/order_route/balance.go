package order_route

import (
	"kcloudb1/internal/handlers/order_handler"
	"kcloudb1/internal/middleware"

	"github.com/gin-gonic/gin"
)

func BalanceRoute(r *gin.RouterGroup) {
	balanceRoute := r.Group("/balance")
	{
		balanceRoute.GET("/", middleware.Auth(), order_handler.GetBalance)
		balanceRoute.GET("/history", middleware.Auth(), order_handler.GetBalanceHistory)
	}
}
