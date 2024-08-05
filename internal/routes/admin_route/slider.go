package admin_route

import (
	"kcloudb1/internal/handlers/admin_handler"
	"kcloudb1/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SliderRoute(r *gin.RouterGroup) {
	sliderRoute := r.Group("/slider")
	{
		sliderRoute.GET("", admin_handler.GetSliderList)
		sliderRoute.POST("", middleware.AuthAdmin(), admin_handler.CreateSlider)
		sliderRoute.PATCH("", middleware.AuthAdmin(), admin_handler.UpdateSlider)
		sliderRoute.DELETE("/by-id", middleware.AuthAdmin(), admin_handler.DeleteSlider)
		sliderRoute.GET("/by-id", admin_handler.GetSlider)
	}
}
