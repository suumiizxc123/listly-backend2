package song_route

import (
	"kcloudb1/internal/handlers/song_handler"

	"github.com/gin-gonic/gin"
)

func SongRoute(r *gin.RouterGroup) {
	songRoute := r.Group("/song")
	{
		songRoute.POST("/", song_handler.CreateSong)
		songRoute.PATCH("/", song_handler.UpdateSong)
		songRoute.GET("/", song_handler.GetSongList)
		songRoute.GET("/by-id", song_handler.GetSong)
		songRoute.DELETE("/by-id", song_handler.DeleteSong)

	}
}
