package artist_route

import (
	"kcloudb1/internal/handlers/artist_handler"

	"github.com/gin-gonic/gin"
)

func ArtistMemberRoute(r *gin.RouterGroup) {
	artistMemberRoute := r.Group("/artist-member")
	{
		artistMemberRoute.POST("/", artist_handler.CreateArtistMember)
		artistMemberRoute.GET("/", artist_handler.GetArtistMemberList)
		artistMemberRoute.GET("/by-id", artist_handler.GetArtistMember)
		artistMemberRoute.DELETE("/by-id", artist_handler.DeleteArtistMember)
		artistMemberRoute.PATCH("/", artist_handler.UpdateArtistMember)
	}
}
