package artist_route

import (
	"kcloudb1/internal/handlers/artist_handler"

	"github.com/gin-gonic/gin"
)

func ArtistBandCombinationRoute(r *gin.RouterGroup) {
	artistBandCombinationRoute := r.Group("/artist-band-combination")
	{
		artistBandCombinationRoute.POST("/", artist_handler.CreateArtistBandCombination)
		artistBandCombinationRoute.GET("/", artist_handler.GetArtistBandCombinationList)
		artistBandCombinationRoute.GET("/by-id", artist_handler.GetArtistBandCombination)
		artistBandCombinationRoute.DELETE("/by-id", artist_handler.DeleteArtistBandCombination)
		artistBandCombinationRoute.PATCH("/", artist_handler.UpdateArtistBandCombination)
	}
}
