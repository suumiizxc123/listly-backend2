package order_route

import (
	"kcloudb1/internal/handlers/order_handler"

	"github.com/gin-gonic/gin"
)

func PaymentRoute(r *gin.RouterGroup) {
	paymentRoute := r.Group("/payment")
	{
		paymentRoute.GET("/:newuid", order_handler.CheckPaymentCallBack)
	}
}
