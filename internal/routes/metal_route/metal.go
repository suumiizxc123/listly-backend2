package metal_route

import (
	"kcloudb1/internal/handlers/metal_handler"

	"github.com/gin-gonic/gin"
)

func MetalRoute(r *gin.RouterGroup) {

	metalRoute := r.Group("/metal")
	{
		metalRoute.GET("/by-id", metal_handler.GetMetal)
		metalRoute.GET("/", metal_handler.GetAllMetals)
		metalRoute.POST("/", metal_handler.CreateMetal)
		metalRoute.PATCH("/", metal_handler.UpdateMetal)
		metalRoute.POST("/by-id", metal_handler.DeleteMetal)
	}
	metalRateRoute := r.Group("/metal-rate")
	{
		metalRateRoute.GET("/last", metal_handler.GetLastMetalRate)
		metalRateRoute.GET("/key", metal_handler.GetMetalRateByKey)
		metalRateRoute.GET("/start-to-end", metal_handler.GetMetalRateByStartToEnd)
		metalRateRoute.POST("/", metal_handler.CreateMetalRate)
	}
}
