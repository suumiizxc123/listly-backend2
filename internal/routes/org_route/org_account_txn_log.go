package org_route

import (
	"kcloudb1/internal/handlers/org_handler"

	"github.com/gin-gonic/gin"
)

func OrgAccountTxnLogRoute(r *gin.RouterGroup) {
	orgAccountTxnLogRoute := r.Group("/org-account-txn-log")
	{
		orgAccountTxnLogRoute.POST("/", org_handler.CreateOrgAccountTxnLog)
		orgAccountTxnLogRoute.GET("/", org_handler.GetOrgAccountTxnLogList)
		orgAccountTxnLogRoute.GET("/by-id", org_handler.GetOrgAccountTxnLog)
		orgAccountTxnLogRoute.DELETE("/by-id", org_handler.DeleteOrgAccountTxnLog)
		orgAccountTxnLogRoute.PATCH("/", org_handler.UpdateOrgAccountTxnLog)
	}
}
