package user_route

import (
	"kcloudb1/internal/handlers/user_handler"
	"kcloudb1/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SysUserRoute(r *gin.RouterGroup) {
	sysUserGroup := r.Group("/sys-user")
	{
		sysUserGroup.POST("/register", middleware.CheckSecret(), user_handler.CreateSysUser)
		sysUserGroup.POST("/login", middleware.CheckSecret(), user_handler.LoginSysUser)
		sysUserGroup.GET("/list", middleware.AuthSysUser(), user_handler.GetSysUserList)
		sysUserGroup.GET("/get", middleware.AuthSysUser(), user_handler.GetSysUser)
		sysUserGroup.PATCH("/update", middleware.AuthSysUser(), user_handler.UpdateSysUser)
		sysUserGroup.PATCH("/update-password", middleware.AuthSysUser(), user_handler.UpdateSysUserPassword)
		sysUserGroup.DELETE("/delete", middleware.AuthSysUser(), user_handler.DeleteSysUser)
	}
}
