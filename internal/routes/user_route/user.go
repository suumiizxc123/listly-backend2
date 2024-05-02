package user_route

import (
	"kcloudb1/internal/handlers/user_handler"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.RouterGroup) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", user_handler.CreateUser)
		userGroup.GET("/list", user_handler.GetUserList)
		userGroup.GET("/get", user_handler.GetUser)
		userGroup.PATCH("/update", user_handler.UpdateUser)
		userGroup.DELETE("/delete", user_handler.DeleteUser)
	}
}
