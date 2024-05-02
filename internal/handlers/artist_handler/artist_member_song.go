package artist_handler

import (
	"kcloudb1/internal/models/artist"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateArtistMemberSong(c *gin.Context) {
	var artistMemberSong artist.ArtistMemberSong

	if err := c.ShouldBindJSON(&artistMemberSong); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist member song fields required", err),
		)
		return
	}

	if err := artistMemberSong.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist member song creation failed", err),
		)
		return
	}

	c.JSON(http.StatusCreated, utils.Success("Artist member song created", artistMemberSong))
}

func GetArtistMemberSongList(c *gin.Context) {

	var artistMemberSong artist.ArtistMemberSong

	list, err := artistMemberSong.GetList()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist member song list failed", err),
		)

		return
	}

	c.JSON(200, utils.Success("Artist member song list", list))
}

func UpdateArtistMemberSong(c *gin.Context) {
	var artistMemberSong artist.ArtistMemberSong

	if err := c.ShouldBindJSON(&artistMemberSong); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist member song fields required", err),
		)
		return
	}

	if err := artistMemberSong.Update(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist member song update failed", err),
		)
		return
	}

	c.JSON(200, utils.Success("Artist member song updated", artistMemberSong))
}

func DeleteArtistMemberSong(c *gin.Context) {
	var artistMemberSong artist.ArtistMemberSong
	var err error
	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist member song id required", nil),
		)
		return
	}

	artistMemberSong.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist member song id cannot be parsed", err),
		)

		return
	}

	if err = artistMemberSong.Delete(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist member song delete failed", err),
		)

		return
	}

	c.JSON(200, utils.Success("Artist member song deleted", artistMemberSong))
}
