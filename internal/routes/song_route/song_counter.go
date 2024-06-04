package song_route

import (
	"kcloudb1/internal/handlers/song_handler"

	"github.com/gin-gonic/gin"
)

func SongCounterRoute(r *gin.RouterGroup) {
	songCounterRoute := r.Group("/song-counter")
	{
		songCounterRoute.POST("/", song_handler.CreateSongCounter)
		songCounterRoute.GET("/", song_handler.GetSongCounterList)
		songCounterRoute.GET("/by-id", song_handler.GetSongCounter)
		songCounterRoute.DELETE("/by-id", song_handler.DeleteSongCounter)
		songCounterRoute.PATCH("/", song_handler.UpdateSongCounter)
	}
}
