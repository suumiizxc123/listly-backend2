package saving_route

import (
	"kcloudb1/internal/handlers/saving_handler"
	"kcloudb1/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SavingRoute(r *gin.RouterGroup) {
	savingRoute := r.Group("/saving")
	{
		savingRoute.POST("", middleware.Auth(), saving_handler.CreateSavingOrder)
		savingRoute.GET("", middleware.Auth(), saving_handler.GetSavingOrder)

	}
}
