package user_route

import (
	"kcloudb1/internal/handlers/user_handler"

	"github.com/gin-gonic/gin"
)

func ServiceLogRoute(r *gin.RouterGroup) {

	serviceLogGroup := r.Group("/service-log")
	{
		serviceLogGroup.POST("/", user_handler.CreateServiceLog)
		serviceLogGroup.GET("/", user_handler.GetServiceLogList)
		serviceLogGroup.PATCH("/", user_handler.UpdateServiceLog)
		serviceLogGroup.DELETE("/", user_handler.DeleteServiceLog)
	}

}
