package song_route

import (
	"kcloudb1/internal/handlers/song_handler"

	"github.com/gin-gonic/gin"
)

func SongCategoryCombinationRoute(r *gin.RouterGroup) {

	songCategoryCombinationRoute := r.Group("/song-category-combination")
	{

		songCategoryCombinationRoute.POST("/", song_handler.CreateSongCategoryCombination)
		songCategoryCombinationRoute.PATCH("/", song_handler.UpdateSongCategoryCombination)
		songCategoryCombinationRoute.GET("/", song_handler.GetSongCategoryCombinationList)
		songCategoryCombinationRoute.GET("/by-id", song_handler.GetSongCategoryCombination)
		songCategoryCombinationRoute.DELETE("/by-id", song_handler.DeleteSongCategoryCombination)
	}
}
