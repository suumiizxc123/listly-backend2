package artist_route

import (
	"kcloudb1/internal/handlers/artist_handler"

	"github.com/gin-gonic/gin"
)

func ArtistProfileRoute(r *gin.RouterGroup) {

	artistProfileRoute := r.Group("/artist-profile")
	{

		artistProfileRoute.POST("/", artist_handler.CreateArtistProfile)
		artistProfileRoute.GET("/", artist_handler.GetArtistProfileList)
		artistProfileRoute.GET("/by-id", artist_handler.GetArtistProfile)
		artistProfileRoute.DELETE("/by-id", artist_handler.DeleteArtistProfile)
		artistProfileRoute.PATCH("/", artist_handler.UpdateArtistProfile)
	}
}
