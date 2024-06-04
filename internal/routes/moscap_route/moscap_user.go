package moscap_route

import (
	moscap_handler "kcloudb1/internal/handlers/mocsap_handler"

	"github.com/gin-gonic/gin"
)

func MoscapUserRoute(r *gin.RouterGroup) {
	moscapUserRoute := r.Group("/moscap-user")
	{
		moscapUserRoute.POST("/", moscap_handler.CreateMoscapUser)
		moscapUserRoute.GET("/", moscap_handler.GetMosCapUserList)
		moscapUserRoute.GET("/by-id", moscap_handler.GetMosCapUser)
		moscapUserRoute.DELETE("/by-id", moscap_handler.DeleteMosCapUser)
		moscapUserRoute.PATCH("/", moscap_handler.UpdateMosCapUser)
	}
}
