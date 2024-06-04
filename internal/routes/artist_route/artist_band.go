package artist_route

import (
	"kcloudb1/internal/handlers/artist_handler"

	"github.com/gin-gonic/gin"
)

func ArtistBandRoute(r *gin.RouterGroup) {
	artistBandRoute := r.Group("/artist-band")
	{
		artistBandRoute.POST("/", artist_handler.CreateArtistBand)
		artistBandRoute.GET("/", artist_handler.GetArtistBandList)
		artistBandRoute.GET("/by-id", artist_handler.GetArtistBand)
		artistBandRoute.DELETE("/by-id", artist_handler.DeleteArtistBand)
		artistBandRoute.PATCH("/", artist_handler.UpdateArtistBand)
	}
}
