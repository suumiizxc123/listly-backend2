package admin_route

import (
	"kcloudb1/internal/handlers/admin_handler"
	"kcloudb1/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoute(r *gin.RouterGroup) {
	adminRoute := r.Group("/admin")
	{
		adminRoute.POST("/login", admin_handler.LoginByPassword)
		adminRoute.GET("/client", middleware.AuthAdmin(), middleware.Paginate(), admin_handler.GetClientList)
		adminRoute.PATCH("/client", middleware.AuthAdmin(), admin_handler.UpdateClient)
		adminRoute.DELETE("/client", middleware.AuthAdmin(), admin_handler.DeleteClient)
		adminRoute.GET("/balance", middleware.AuthAdmin(), admin_handler.GetBalanceByClientID)
		adminRoute.GET("/order", middleware.AuthAdmin(), middleware.Paginate(), admin_handler.GetOrderList)

		adminRoute.POST("/order/verify", middleware.AuthAdmin(), admin_handler.VerifyOrder)
		adminRoute.POST("/order/cancel", middleware.AuthAdmin(), admin_handler.CancelOrder)
		adminRoute.POST("/order/create", middleware.AuthAdmin(), admin_handler.CreateOrder)

		adminRoute.POST("/withdraw/verify", middleware.AuthAdmin(), admin_handler.VerifyWithDraw)
		adminRoute.POST("/withdraw", middleware.AuthAdmin(), admin_handler.CancelOrder)

		adminRoute.GET("/withdraw", middleware.AuthAdmin(), middleware.Paginate(), admin_handler.GetWithDrawList)

		adminRoute.POST("/message", middleware.AuthAdmin(), admin_handler.SendMessage)
		adminRoute.POST("/saving/verify", middleware.AuthAdmin(), admin_handler.VerifySaving)

		// adminRoute.GET("/faq", middleware.AuthAdmin(), admin_handler.GetFAQList)
		// adminRoute.POST("/faq", middleware.AuthAdmin(), admin_handler.CreateFAQ)
		// adminRoute.PATCH("/faq", middleware.AuthAdmin(), admin_handler.UpdateFAQ)
		// adminRoute.DELETE("/faq/by-id", middleware.AuthAdmin(), admin_handler.DeleteFAQ)
		// adminRoute.GET("/faq/by-id", middleware.AuthAdmin(), admin_handler.GetFAQ)
	}

}
