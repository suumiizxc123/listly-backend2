package admin_route

import (
	"kcloudb1/internal/handlers/admin_handler"
	"kcloudb1/internal/middleware"

	"github.com/gin-gonic/gin"
)

func IngredientRoute(r *gin.RouterGroup) {

	ingredientRoute := r.Group("/ingredient")
	{
		ingredientRoute.GET("/", admin_handler.GetIngredientList)
		ingredientRoute.POST("/", middleware.AuthAdmin(), admin_handler.CreateIngredient)
		ingredientRoute.PATCH("/", middleware.AuthAdmin(), admin_handler.UpdateIngredient)
		ingredientRoute.DELETE("/by-id", middleware.AuthAdmin(), admin_handler.DeleteIngredient)
		ingredientRoute.GET("/by-id", admin_handler.GetIngredient)
	}
}
