package user_route

import (
	"kcloudb1/internal/handlers/user_handler"

	"github.com/gin-gonic/gin"
)

func SysUserRoute(r *gin.RouterGroup) {
	sysUserGroup := r.Group("/sys-user")
	{
		sysUserGroup.POST("/register", user_handler.CreateSysUser)
		sysUserGroup.POST("/login", user_handler.LoginSysUser)
		sysUserGroup.GET("/list", user_handler.GetSysUserList)
		sysUserGroup.GET("/get", user_handler.GetSysUser)
		sysUserGroup.PATCH("/update", user_handler.UpdateSysUser)
		sysUserGroup.DELETE("/delete", user_handler.DeleteSysUser)
	}
}
