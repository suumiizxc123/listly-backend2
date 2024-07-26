package admin_route

import (
	"kcloudb1/internal/handlers/admin_handler"
	"kcloudb1/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewsRoute(r *gin.RouterGroup) {

	newsRoute := r.Group("/news")
	{
		newsRoute.GET("/", admin_handler.GetNewsList)
		newsRoute.POST("/", middleware.AuthAdmin(), admin_handler.CreateNews)
		newsRoute.PATCH("/", middleware.AuthAdmin(), admin_handler.UpdateNews)
		newsRoute.DELETE("/by-id", middleware.AuthAdmin(), admin_handler.DeleteNews)
		newsRoute.GET("/by-id", admin_handler.GetNews)
	}
}
