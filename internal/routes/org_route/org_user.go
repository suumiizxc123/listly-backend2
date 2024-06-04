package org_route

import (
	"kcloudb1/internal/handlers/org_handler"

	"github.com/gin-gonic/gin"
)

func OrgUserRoute(r *gin.RouterGroup) {
	orgUserRoute := r.Group("/org-user")
	{
		orgUserRoute.POST("/", org_handler.CreateOrgUser)
		orgUserRoute.GET("/", org_handler.GetOrgUserList)
		orgUserRoute.GET("/by-id", org_handler.GetOrgUser)
		orgUserRoute.DELETE("/by-id", org_handler.DeleteOrgUser)
		orgUserRoute.PATCH("/", org_handler.UpdateOrgUser)
	}
}
