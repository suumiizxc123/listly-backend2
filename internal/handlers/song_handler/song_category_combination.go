package song_handler

import (
	"kcloudb1/internal/models/song"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateSongCategoryCombination(c *gin.Context) {

	var songCategoryCombination song.SongCategoryCombination

	if err := c.ShouldBindJSON(&songCategoryCombination); err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Song category combination fields required", err),
		)

		return
	}

	if err := songCategoryCombination.Create(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Song category combination creation failed", err),
		)

		return
	}

	c.JSON(http.StatusCreated, utils.Success("Song category combination created", songCategoryCombination))

}

func GetSongCategoryCombinationList(c *gin.Context) {

	var songCategoryCombination song.SongCategoryCombination

	songCategoryCombinations, err := songCategoryCombination.GetList()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Song category combination list retrieval failed", err),
		)

		return
	}

	c.JSON(200, utils.Success("Song category combination list", songCategoryCombinations))
}

func GetSongCategoryCombination(c *gin.Context) {

	var songCategoryCombination song.SongCategoryCombination
	var err error
	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Song category combination id required", "id must be required"),
		)

		return
	}

	songCategoryCombination.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Song category combination id cannot be parsed", err.Error()),
		)

		return
	}

	if err = songCategoryCombination.Get(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Song category combination retrieval failed", err),
		)

		return
	}

	c.JSON(200, utils.Success("Song category combination", songCategoryCombination))

}

func UpdateSongCategoryCombination(c *gin.Context) {

	var songCategoryCombination song.SongCategoryCombination

	if err := c.ShouldBindJSON(&songCategoryCombination); err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Song category combination fields required", err),
		)

		return
	}

	if err := songCategoryCombination.Update(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Song category combination update failed", err),
		)

		return
	}

	c.JSON(200, utils.Success("Song category combination updated", songCategoryCombination))
}

func DeleteSongCategoryCombination(c *gin.Context) {

	var songCategoryCombination song.SongCategoryCombination
	var err error
	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Song category combination id required", "id must be required"),
		)

		return
	}

	songCategoryCombination.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Song category combination id cannot be parsed", err.Error()),
		)

		return
	}

	if err = songCategoryCombination.Delete(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Song category combination delete failed", err),
		)

		return
	}

	c.JSON(200, utils.Success("Song category combination deleted", songCategoryCombination))
}
