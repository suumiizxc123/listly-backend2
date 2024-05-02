package common_route

import (
	"kcloudb1/internal/handlers/common_handler"

	"github.com/gin-gonic/gin"
)

func LanguageRoute(r *gin.RouterGroup) {
	languageRoute := r.Group("/language")
	{
		languageRoute.POST("/", common_handler.CreateLanguage)
		languageRoute.GET("/", common_handler.GetLanguageList)
		languageRoute.PATCH("/", common_handler.UpdateLanguage)
		languageRoute.DELETE("/", common_handler.DeleteLanguage)
	}
}
