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
			utils.Error("Artist profile fields required", err),
		)
		return
	}

	if err := artistProfile.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist profile creation failed", err),
		)
		return
	}

	c.JSON(http.StatusCreated, utils.Success("Artist profile created", artistProfile))
}

func GetArtistProfile(c *gin.Context) {
	var artistProfile artist.ArtistProfile
	var err error
	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist profile id required", nil),
		)
		return
	}

	artistProfile.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist profile id cannot be parsed", err),
		)
		return
	}

	if err = artistProfile.Get(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist profile not found", err),
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
			utils.Error("Artist profile list not found", err),
		)

		return
	}

	c.JSON(200, utils.Success("Artist profile list", list))
}

func UpdateArtistProfile(c *gin.Context) {
	var artistProfile artist.ArtistProfile

	if err := c.ShouldBind(&artistProfile); err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist profile fields required", err),
		)

		return
	}

	if err := artistProfile.Update(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist profile update failed", err),
		)

		return
	}

	c.JSON(200, utils.Success("Artist profile updated", artistProfile))
}

func DeleteArtistProfile(c *gin.Context) {

	var artistProfile artist.ArtistProfile
	var err error
	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist profile id required", nil),
		)
		return
	}

	artistProfile.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist profile id cannot be parsed", err),
		)
		return

	}

	if err = artistProfile.Delete(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist profile delete failed", err),
		)
		return
	}

	c.JSON(200, utils.Success("Artist profile deleted", artistProfile))
}
