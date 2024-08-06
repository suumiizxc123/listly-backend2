package order_route

import (
	"kcloudb1/internal/handlers/order_handler"

	"github.com/gin-gonic/gin"
)

func PaymentRoute(r *gin.RouterGroup) {
	paymentRoute := r.Group("/payment")
	{
		paymentRoute.GET("/vip-member/:newuid", order_handler.CheckPaymentVIPMemberCallBack)
		paymentRoute.GET("/:newuid", order_handler.CheckPaymentCallBack)
		paymentRoute.GET("/saving/:newuid", order_handler.CheckSavingPaymentCallBack)
		paymentRoute.POST("/check-payment/:newuid", order_handler.CheckPayment)
	}

}
