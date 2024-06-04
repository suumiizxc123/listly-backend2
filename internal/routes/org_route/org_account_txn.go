package org_route

import (
	"kcloudb1/internal/handlers/org_handler"

	"github.com/gin-gonic/gin"
)

func OrgAccountTxnRoute(r *gin.RouterGroup) {
	orgAccountTxnRoute := r.Group("/org-account-txn")
	{
		orgAccountTxnRoute.POST("/", org_handler.CreateOrgAccountTxn)
		orgAccountTxnRoute.GET("/", org_handler.GetOrgAccountTxnList)
		orgAccountTxnRoute.GET("/by-id", org_handler.GetOrgAccountTxn)
		orgAccountTxnRoute.DELETE("/by-id", org_handler.DeleteOrgAccountTxn)
		orgAccountTxnRoute.PATCH("/", org_handler.UpdateOrgAccountTxn)
	}
}
