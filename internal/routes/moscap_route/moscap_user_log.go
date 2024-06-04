package moscap_route

import (
	moscap_handler "kcloudb1/internal/handlers/mocsap_handler"

	"github.com/gin-gonic/gin"
)

func MoscapUserLogRoute(r *gin.RouterGroup) {
	moscapUserLogRoute := r.Group("/moscap-user-log")
	{
		moscapUserLogRoute.POST("/", moscap_handler.CreateMoscapUserLog)
		moscapUserLogRoute.GET("/", moscap_handler.GetMosCapUserLogList)
		moscapUserLogRoute.GET("/by-id", moscap_handler.GetMosCapUserLog)
		moscapUserLogRoute.DELETE("/by-id", moscap_handler.DeleteMosCapUserLog)
		moscapUserLogRoute.PATCH("/", moscap_handler.UpdateMosCapUserLog)
	}
}
