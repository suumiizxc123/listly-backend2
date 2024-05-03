package song_handler

import (
	"kcloudb1/internal/models/song"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateSong(c *gin.Context) {
	var song song.Song
	var err error

	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("Song fields required", err),
		)
		return
	}

	song.UUID = uuid.New().String()

	if err = song.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Song creation failed", err),
		)
		return
	}

	c.JSON(http.StatusCreated, utils.Success("Song created", song))
}

func UpdateSong(c *gin.Context) {
	var song song.Song
	var err error

	if err := c.ShouldBindJSON(&song); err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Song fields required", err),
		)
		return
	}

	if err = song.Update(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Song update failed", err),
		)
		return
	}

	c.JSON(200, utils.Success("Song updated", song))

}

func GetSong(c *gin.Context) {

	var song song.SongExtend
	var err error
	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Song id required", "id must be required"),
		)

		return
	}

	song.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Song id cannot be parsed", err.Error()),
		)

		return
	}

	if err = song.Get(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Song not found", err.Error()),
		)

		return
	}

	c.JSON(200, utils.Success("Song found", song))
}

func DeleteSong(c *gin.Context) {

	var song song.Song
	var err error

	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Song id required", "id must be required"),
		)

		return
	}

	song.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Song id cannot be parsed", err.Error()),
		)

		return
	}

	if err = song.Delete(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Song deletion failed", err.Error()),
		)

		return
	}

	c.JSON(200, utils.Success("Song deleted", struct{}{}))
}

func GetSongList(c *gin.Context) {
	var song song.SongExtend

	songs, err := song.GetList()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Song list not found", err.Error()),
		)

		return
	}

	c.JSON(200, utils.Success("Song list", songs))
}
