package song_route

import (
	"kcloudb1/internal/handlers/song_handler"

	"github.com/gin-gonic/gin"
)

func SongCategoryRoute(r *gin.RouterGroup) {
	songCategoryRoute := r.Group("/song-category")
	{
		songCategoryRoute.POST("/", song_handler.CreateSongCategory)
		songCategoryRoute.PATCH("/", song_handler.UpdateSongCategory)
		songCategoryRoute.GET("/", song_handler.GetSongCategoryList)
		songCategoryRoute.GET("/by-id", song_handler.GetSongCategory)
		songCategoryRoute.DELETE("/by-id", song_handler.DeleteSongCategory)
	}
}
