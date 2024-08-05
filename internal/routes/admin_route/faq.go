package admin_route

import (
	"kcloudb1/internal/handlers/admin_handler"
	"kcloudb1/internal/middleware"

	"github.com/gin-gonic/gin"
)

func FAQRoute(r *gin.RouterGroup) {

	faqRoute := r.Group("/faq")
	{
		faqRoute.GET("/", middleware.AuthAdmin(), admin_handler.GetFAQList)
		faqRoute.POST("/", middleware.AuthAdmin(), admin_handler.CreateFAQ)
		faqRoute.PATCH("/", middleware.AuthAdmin(), admin_handler.UpdateFAQ)
		faqRoute.DELETE("/by-id", middleware.AuthAdmin(), admin_handler.DeleteFAQ)
		faqRoute.GET("/by-id", middleware.AuthAdmin(), admin_handler.GetFAQ)
	}
}
