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
			utils.Error([]string{"Artist song fields required", "Артист дууны мэдээлэл дутуу байна"}, err),
		)
		return
	}

	if err := artistSong.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist song creation failed", "Артист дууг хадгалахад алдаа үүслээ"}, err),
		)
		return
	}

	c.JSON(http.StatusCreated, utils.Success([]string{"Artist song created", "Артист дууг хадгаллаа"}, artistSong))
}

func GetArtistSongList(c *gin.Context) {
	var artistSong artist.ArtistSong

	list, err := artistSong.GetList()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist song list not found", "Артист дуугуудын жагсаалт олдсонгүй"}, err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"Artist song list", "Артист дуугуудын жагсаалт татлаа"}, list))
}

func UpdateArtistSong(c *gin.Context) {
	var artistSong artist.ArtistSong

	if err := c.ShouldBindJSON(&artistSong); err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist song fields required", "Артист дууны мэдээлэл дутуу байна"}, err),
		)
		return
	}

	if err := artistSong.Update(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist song update failed", "Артист дууг хадгалахад алдаа үүслээ"}, err),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"Artist song updated", "Артист дууг хадгаллаа"}, artistSong))
}

func DeleteArtistSong(c *gin.Context) {
	var artistSong artist.ArtistSong
	var err error
	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist song id required", "Артист дууны id бөглөнө үү"}, nil),
		)
		return
	}

	artistSong.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist song id cannot be parsed", "Артист дууны id буруу байна"}, err),
		)
		return
	}

	if err = artistSong.Delete(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist song delete failed", "Артист дууг устгаж чадсангүй"}, err),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"Artist song deleted", "Артист дууг устгалаа"}, artistSong))
}

func GetArtistSong(c *gin.Context) {
	var artistSong artist.ArtistSong
	var err error
	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist song id required", "Артист дууны id бөглөнө үү"}, nil),
		)
		return
	}

	artistSong.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist song id cannot be parsed", "Артист дууны id буруу байна"}, err),
		)
		return
	}

	if err = artistSong.Get(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist song not found", "Артист дууг олсонгүй"}, err),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"Artist song", "Артист дууг татлаа"}, artistSong))
}
