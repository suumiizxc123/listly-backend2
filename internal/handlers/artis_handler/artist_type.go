package artist_handler

import (
	"kcloudb1/internal/models/artist"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateArtistType(c *gin.Context) {
	var artisType artist.ArtistType

	if err := c.ShouldBindJSON(&artisType); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist type fields required", err),
		)
		return
	}

	if err := artisType.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist type creation failed", err),
		)
		return
	}

	c.JSON(http.StatusCreated, utils.Success("Artist type created", artisType))
}

func GetArtistTypeList(c *gin.Context) {
	var artisType artist.ArtistType

	list, err := artisType.GetList()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist type list failed", err),
		)
		return
	}

	c.JSON(200, utils.Success("Artist type list", list))
}

func GetArtistTypeListByType(c *gin.Context) {
	var artisType artist.ArtistType

	tp, ok := c.GetQuery("type")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist type type required", nil),
		)
		return
	}

	artisType.Type = tp

	list, err := artisType.GetListByType()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist type list failed", err),
		)
		return
	}

	c.JSON(200, utils.Success("Artist type list", list))
}

func GetArtistType(c *gin.Context) {

	var artisType artist.ArtistType
	var err error
	tp, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist type id required", nil),
		)

		return
	}

	artisType.ID, err = strconv.ParseInt(tp, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist type id cannot be parsed", err),
		)

		return
	}

	if err = artisType.Get(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist type not found", err),
		)

		return
	}

	c.JSON(200, utils.Success("Artist type found", artisType))
}

func UpdateArtistType(c *gin.Context) {

	var artisType artist.ArtistType

	if err := c.ShouldBindJSON(&artisType); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist type fields required", err),
		)
		return
	}

	if err := artisType.Update(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist type update failed", err),
		)

		return
	}

	c.JSON(200, utils.Success("Artist type updated", artisType))

}

func DeleteArtistType(c *gin.Context) {

	var artisType artist.ArtistType
	var err error
	tp, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist type id required", nil),
		)

		return
	}

	artisType.ID, err = strconv.ParseInt(tp, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Artist type id cannot be parsed", err),
		)

		return
	}

	if err = artisType.Delete(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Artist type delete failed", err),
		)

		return
	}

	c.JSON(200, utils.Success("Artist type deleted", struct{}{}))
}
