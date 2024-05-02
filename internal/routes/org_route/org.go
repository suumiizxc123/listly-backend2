package org_route

import (
	"kcloudb1/internal/handlers/org_handler"

	"github.com/gin-gonic/gin"
)

func OrgRoute(r *gin.RouterGroup) {
	orgRoute := r.Group("/org")
	{
		orgRoute.POST("/", org_handler.CreateOrg)
		orgRoute.PATCH("/", org_handler.UpdateOrg)
		orgRoute.GET("/", org_handler.GetOrgList)
		orgRoute.GET("/by-id", org_handler.GetOrg)
		orgRoute.DELETE("/by-id", org_handler.DeleteOrg)
	}
}
