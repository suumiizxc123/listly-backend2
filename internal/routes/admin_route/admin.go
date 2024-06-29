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
	}
}
