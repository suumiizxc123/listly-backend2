package artist_handler

import (
	"kcloudb1/internal/models/artist"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateArtistMember(c *gin.Context) {
	var artistMember artist.ArtistMember

	if err := c.ShouldBind(&artistMember); err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist member fields required", "Артист груп дууны мэдээлэл үүсгэхдээ бүрэн бөглөнө үү"}, err),
		)

		return
	}

	if err := artistMember.Create(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist profile creation failed", "Артист груп үүсгэхэд алдаа гарлаа"}, err),
		)

		return
	}

	c.JSON(http.StatusCreated, utils.Success([]string{"Artist profile created", "Артист груп үүсгэгдсэн"}, artistMember))

}

func GetArtistMemberList(c *gin.Context) {

	var artistMember artist.ArtistMember

	list, err := artistMember.GetList()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist profile list not found", "Артист групууд олдсонгүй"}, err),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"Artist profile list", "Артист групууд лист татлаа"}, list))
}

func UpdateArtistMember(c *gin.Context) {

	var artistMember artist.ArtistMember

	if err := c.ShouldBind(&artistMember); err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist member fields required", "Артист груп дууны мэдээлэл үүсгэхдээ бүрэн бөглөнө үү"}, err),
		)

		return
	}

	if err := artistMember.Update(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist profile update failed", "Артист хадгалахад алдаа гарлаа"}, err),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"Artist profile updated", "Артист хадгаллаа"}, artistMember))
}

func DeleteArtistMember(c *gin.Context) {

	var artistMember artist.ArtistMember
	var err error
	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist profile id required", "Артист груп дууны id бөглөнө үү"}, nil),
		)

		return
	}

	artistMember.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist profile id cannot be parsed", "Артист груп дууны id буруу байна"}, err),
		)

		return
	}

	if err = artistMember.Delete(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist profile deletion failed", "Артист груп устгахад алдаа гарлаа"}, err),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"Artist profile deleted", "Артист груп устгалаа"}, nil))
}

func GetArtistMember(c *gin.Context) {

	var artistMember artist.ArtistMember
	var err error
	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist profile id required", "Артист груп дууны id бөглөнө үү"}, nil),
		)

		return
	}

	artistMember.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist profile id cannot be parsed", "Артист груп дууны id буруу байна"}, err),
		)

		return
	}

	if err = artistMember.Get(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist profile not found", "Артист олдсонгүй"}, err),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"Artist profile found", "Артист олдсон"}, artistMember))
}
