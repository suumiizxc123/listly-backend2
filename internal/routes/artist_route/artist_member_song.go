package artist_route

import (
	"kcloudb1/internal/handlers/artist_handler"

	"github.com/gin-gonic/gin"
)

func ArtistMemberSongRoute(r *gin.RouterGroup) {

	artistMemberSongRoute := r.Group("/artist-member-song")
	{

		artistMemberSongRoute.POST("/", artist_handler.CreateArtistMemberSong)
		artistMemberSongRoute.GET("/", artist_handler.GetArtistMemberSongList)
		artistMemberSongRoute.DELETE("/by-id", artist_handler.DeleteArtistMemberSong)
		artistMemberSongRoute.PATCH("/", artist_handler.UpdateArtistMemberSong)

	}
}
