package upload_route

import (
	"kcloudb1/internal/handlers/upload_handler"

	"github.com/gin-gonic/gin"
)

func UploadRoute(r *gin.RouterGroup) {
	uploadRoute := r.Group("/upload")
	{
		uploadRoute.POST("/image", upload_handler.UploadImage)
	}
}
