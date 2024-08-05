package admin_route

import (
	"kcloudb1/internal/handlers/admin_handler"
	"kcloudb1/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoute(r *gin.RouterGroup) {
	productRoute := r.Group("/product")
	{
		productRoute.GET("", admin_handler.GetProductList)
		productRoute.POST("", middleware.AuthAdmin(), admin_handler.CreateProduct)
		productRoute.PATCH("", middleware.AuthAdmin(), admin_handler.UpdateProduct)
		productRoute.DELETE("/by-id", middleware.AuthAdmin(), admin_handler.DeleteProduct)
		productRoute.GET("/by-id", admin_handler.GetProduct)

		productRoute.POST("/image", middleware.AuthAdmin(), admin_handler.AddProductImage)
		productRoute.DELETE("/image", middleware.AuthAdmin(), admin_handler.RemoveProductImage)

		productRoute.POST("/ingredient", middleware.AuthAdmin(), admin_handler.AddProductIngredient)
		productRoute.DELETE("/ingredient", middleware.AuthAdmin(), admin_handler.RemoveProductIngredient)
	}
}
