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
			utils.Error("Artist profile fields required", err),
		)

		return
	}

	if err := artistMember.Create(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist profile creation failed", err),
		)

		return
	}

	c.JSON(http.StatusCreated, utils.Success("Artist profile created", artistMember))

}

func GetArtistMemberList(c *gin.Context) {

	var artistMember artist.ArtistMember

	list, err := artistMember.GetList()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist profile list not found", err),
		)
		return
	}

	c.JSON(200, utils.Success("Artist profile list", list))
}

func UpdateArtistMember(c *gin.Context) {

	var artistMember artist.ArtistMember

	if err := c.ShouldBind(&artistMember); err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist profile fields required", err),
		)

		return
	}

	if err := artistMember.Update(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist profile update failed", err),
		)

		return
	}

	c.JSON(200, utils.Success("Artist profile updated", artistMember))
}

func DeleteArtistMember(c *gin.Context) {

	var artistMember artist.ArtistMember
	var err error
	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist profile id required", nil),
		)

		return
	}

	artistMember.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist profile id cannot be parsed", err),
		)

		return
	}

	if err = artistMember.Delete(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist profile delete failed", err),
		)

		return
	}

	c.JSON(200, utils.Success("Artist profile deleted", nil))
}

func GetArtistMember(c *gin.Context) {

	var artistMember artist.ArtistMember
	var err error
	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist profile id required", nil),
		)

		return
	}

	artistMember.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist profile id cannot be parsed", err),
		)

		return
	}

	if err = artistMember.Get(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist profile not found", err),
		)

		return
	}

	c.JSON(200, utils.Success("Artist profile found", artistMember))
}
