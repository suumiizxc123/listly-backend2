package artist_route

import (
	"kcloudb1/internal/handlers/artist_handler"

	"github.com/gin-gonic/gin"
)

func ArtistTypeRoute(r *gin.RouterGroup) {
	artistTypeRoute := r.Group("/artist-type")
	{
		artistTypeRoute.POST("/", artist_handler.CreateArtistType)
		artistTypeRoute.GET("/", artist_handler.GetArtistTypeList)
		artistTypeRoute.GET("/by-id", artist_handler.GetArtistType)
		artistTypeRoute.DELETE("/by-id", artist_handler.DeleteArtistType)
		artistTypeRoute.PATCH("/", artist_handler.UpdateArtistType)
		artistTypeRoute.GET("/by-type", artist_handler.GetArtistTypeListByType)
	}
}
