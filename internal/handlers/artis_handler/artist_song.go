package artist_handler

import (
	"kcloudb1/internal/models/artist"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateArtistSong(c *gin.Context) {
	var artistSong artist.ArtistSong

	if err := c.ShouldBindJSON(&artistSong); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist song fields required", err),
		)
		return
	}

	if err := artistSong.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist song creation failed", err),
		)
		return
	}

	c.JSON(http.StatusCreated, utils.Success("Artist song created", artistSong))
}

func GetArtistSongList(c *gin.Context) {
	var artistSong artist.ArtistSong

	list, err := artistSong.GetList()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist song list not found", err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success("Artist song list", list))
}

func UpdateArtistSong(c *gin.Context) {
	var artistSong artist.ArtistSong

	if err := c.ShouldBindJSON(&artistSong); err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist song fields required", err),
		)
		return
	}

	if err := artistSong.Update(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist song update failed", err),
		)
		return
	}

	c.JSON(200, utils.Success("Artist song updated", artistSong))
}

func DeleteArtistSong(c *gin.Context) {
	var artistSong artist.ArtistSong
	var err error
	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist song id required", nil),
		)
		return
	}

	artistSong.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist song id cannot be parsed", err),
		)
		return
	}

	if err = artistSong.Delete(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist song delete failed", err),
		)
		return
	}

	c.JSON(200, utils.Success("Artist song deleted", artistSong))
}

func GetArtistSong(c *gin.Context) {
	var artistSong artist.ArtistSong
	var err error
	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist song id required", nil),
		)
		return
	}

	artistSong.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist song id cannot be parsed", err),
		)
		return
	}

	if err = artistSong.Get(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist song not found", err),
		)
		return
	}

	c.JSON(200, utils.Success("Artist song", artistSong))
}
