package org_route

import (
	"kcloudb1/internal/handlers/org_handler"

	"github.com/gin-gonic/gin"
)

func OrgAccountRoute(r *gin.RouterGroup) {
	orgAccountRoute := r.Group("/org-account")
	{
		orgAccountRoute.POST("/", org_handler.CreateOrgAccount)
		orgAccountRoute.GET("/", org_handler.GetOrgAccountList)
		orgAccountRoute.GET("/by-id", org_handler.GetOrgAccount)
		orgAccountRoute.DELETE("/by-id", org_handler.DeleteOrgAccount)
		orgAccountRoute.PATCH("/", org_handler.UpdateOrgAccount)
	}
}
