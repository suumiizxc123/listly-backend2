package org_route

import (
	"kcloudb1/internal/handlers/org_handler"
	"kcloudb1/internal/middleware"

	"github.com/gin-gonic/gin"
)

func OrgSysRoute(r *gin.RouterGroup) {
	orgSysRoute := r.Group("/sys-user/org")
	{
		orgSysRoute.POST("/and-user", middleware.AuthSysUser(), org_handler.CreateOrgAndUser)
		orgSysRoute.PATCH("/", org_handler.UpdateOrg)
		orgSysRoute.GET("/", middleware.AuthSysUser(), middleware.Paginate(), org_handler.GetOrgList)
		orgSysRoute.GET("/by-id", middleware.AuthSysUser(), org_handler.GetOrg)

		orgSysRoute.DELETE("/", org_handler.DeleteOrg)
	}
}

func OrgRoute(r *gin.RouterGroup) {
	orgRoute := r.Group("/user/org")
	{
		orgRoute.GET("/", middleware.AuthUser(), org_handler.GetOrgListUser)
	}
}
