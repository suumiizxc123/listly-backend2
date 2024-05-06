package artist_handler

import (
	"kcloudb1/internal/models/artist"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateArtistProfile(c *gin.Context) {
	var artistProfile artist.ArtistProfile

	if err := c.ShouldBind(&artistProfile); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist profile fields required", "Артист үүсгэхдээ бүрэн бөглөнө үү"}, err),
		)
		return
	}

	if err := artistProfile.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist profile creation failed", "Артист үүсгэхэд алдаа гарлаа"}, err),
		)
		return
	}

	c.JSON(http.StatusCreated, utils.Success([]string{"Artist profile created", "Артист амжилттай хадгаллаа"}, artistProfile))
}

func GetArtistProfile(c *gin.Context) {
	var artistProfile artist.ArtistProfile
	var err error
	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist profile id required", "Артист сонгогдоогүй байна"}, nil),
		)
		return
	}

	artistProfile.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist profile id cannot be parsed", "Артист id буруу байна"}, err),
		)
		return
	}

	if err = artistProfile.Get(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist profile not found", "Артист олдсонгүй"}, err),
		)
		return
	}

}

func GetArtistProfileList(c *gin.Context) {

	var artistProfile artist.ArtistProfile

	list, err := artistProfile.GetList()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist profile list not found", "Артист олдсонгүй"}, err),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"Artist profile list", "Артист лист татлаа"}, list))
}

func UpdateArtistProfile(c *gin.Context) {
	var artistProfile artist.ArtistProfile

	if err := c.ShouldBind(&artistProfile); err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist profile fields required", "Артист бүтцийн дагуу бүрэн бөглөнө үү"}, err),
		)

		return
	}

	if err := artistProfile.Update(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist profile update failed", "Артист хадгалахад алдаа гарлаа"}, err),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"Artist profile updated", "Артист хадгаллаа"}, artistProfile))
}

func DeleteArtistProfile(c *gin.Context) {

	var artistProfile artist.ArtistProfile
	var err error
	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist profile id required", "Артист сонгогдоогүй байна"}, nil),
		)
		return
	}

	artistProfile.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist profile id cannot be parsed", "Артист id буруу байна"}, err),
		)
		return

	}

	if err = artistProfile.Delete(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist profile delete failed", "Артист устгахад алдаа гарлаа"}, err),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"Artist profile deleted", "Артист устгалаа"}, artistProfile))
}
