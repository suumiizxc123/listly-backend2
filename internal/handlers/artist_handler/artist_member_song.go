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
			utils.Error([]string{"Artist member song fields required", "Артист групийн дууны мэдээлэл бүрэн хангагдаагүй байна"}, err),
		)
		return
	}

	if err := artistMemberSong.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist member song creation failed", "Артист групийн дууг хадгалахад алдаа үүслээ"}, err),
		)
		return
	}

	c.JSON(http.StatusCreated, utils.Success([]string{"Artist member song created", "Артист групийн дууны мэдээлэл амжилттай хадгаллаа"}, artistMemberSong))
}

func GetArtistMemberSongList(c *gin.Context) {

	var artistMemberSong artist.ArtistMemberSong

	list, err := artistMemberSong.GetList()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist member song list failed", "Артист групийн дууны мэдээлэл жагсаалт авахад алдаа гарлаа"}, err),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"Artist member song list", "Артист групийн дууны мэдээлэл жагсаалт авлаа"}, list))
}

func UpdateArtistMemberSong(c *gin.Context) {
	var artistMemberSong artist.ArtistMemberSong

	if err := c.ShouldBindJSON(&artistMemberSong); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist member song fields required", "Артист групийн дууны мэдээлэл бүрэн хангагдаагүй байна"}, err),
		)
		return
	}

	if err := artistMemberSong.Update(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist member song update failed", "Артист групийн дууг хадгалахад алдаа үүслээ"}, err),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"Artist member song updated", "Артист групийн дууны мэдээлэл амжилттай хадгаллаа"}, artistMemberSong))
}

func DeleteArtistMemberSong(c *gin.Context) {
	var artistMemberSong artist.ArtistMemberSong
	var err error
	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist member song id required", "Артист групийн дууны id шаардлагатай байна"}, nil),
		)
		return
	}

	artistMemberSong.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist member song id cannot be parsed", "Артист групийн дууны id таарахгүй байна"}, err),
		)

		return
	}

	if err = artistMemberSong.Delete(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist member song deletion failed", "Артист групийн дууг устгахад алдаа үүслээ"}, err),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"Artist member song deleted", "Артист групийн дууны мэдээлэл амжилттай устгалаа"}, artistMemberSong))
}
