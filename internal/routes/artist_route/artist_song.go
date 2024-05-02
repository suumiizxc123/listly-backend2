package artist_route

import (
	"kcloudb1/internal/handlers/artist_handler"

	"github.com/gin-gonic/gin"
)

func ArtistSongRoute(r *gin.RouterGroup) {
	artistSongRoute := r.Group("/artist-song")
	{
		artistSongRoute.POST("/", artist_handler.CreateArtistSong)
		artistSongRoute.GET("/", artist_handler.GetArtistSongList)
		artistSongRoute.GET("/by-id", artist_handler.GetArtistSong)
		artistSongRoute.DELETE("/by-id", artist_handler.DeleteArtistSong)
		artistSongRoute.PATCH("/", artist_handler.UpdateArtistSong)
	}
}
