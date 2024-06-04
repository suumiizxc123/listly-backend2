package song_route

import (
	"kcloudb1/internal/handlers/song_handler"

	"github.com/gin-gonic/gin"
)

func SongLanguageRoute(r *gin.RouterGroup) {
	songLanguageRoute := r.Group("/song-language")
	{
		songLanguageRoute.POST("/", song_handler.CreateSongLanguage)
		songLanguageRoute.GET("/", song_handler.GetSongLanguageList)
		songLanguageRoute.GET("/by-id", song_handler.GetSongLanguage)
		songLanguageRoute.DELETE("/by-id", song_handler.DeleteSongLanguage)
		songLanguageRoute.PATCH("/", song_handler.UpdateSongLanguage)
	}
}
