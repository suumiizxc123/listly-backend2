package user_route

import (
	"kcloudb1/internal/handlers/user_handler"
	"kcloudb1/internal/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.RouterGroup) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", middleware.CheckSecret(), user_handler.CreateUser)
		userGroup.POST("/login", middleware.CheckSecret(), user_handler.LoginUser)
		userGroup.GET("/list", user_handler.GetUserList)
		userGroup.GET("/get", middleware.AuthUser(), user_handler.GetUser)
		userGroup.PATCH("/update", middleware.AuthUser(), user_handler.UpdateUser)
		userGroup.PATCH("/update-password", middleware.AuthUser(), user_handler.UpdateUserPassword)
		userGroup.DELETE("/delete", middleware.AuthUser(), user_handler.DeleteUser)
		//update unique phone number
		// test
	}
}
