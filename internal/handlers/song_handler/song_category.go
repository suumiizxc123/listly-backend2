package song_handler

import (
	"kcloudb1/internal/models/song"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateSongCategory(c *gin.Context) {
	var songCategory song.SongCategory
	var err error

	if err := c.ShouldBindJSON(&songCategory); err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Song category fields required", err),
		)

		return
	}

	if err = songCategory.Create(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Song category creation failed", err),
		)

		return
	}

	c.JSON(http.StatusCreated, utils.Success("Song category created", songCategory))
}

func GetSongCategoryList(c *gin.Context) {

	var songCategory song.SongCategory

	songCategories, err := songCategory.GetList()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Song category list not found", err.Error()),
		)

		return
	}

	c.JSON(200, utils.Success("Song category list", songCategories))
}

func GetSongCategory(c *gin.Context) {

	var songCategory song.SongCategory
	var err error
	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Song category id required", "id must be required"),
		)

		return

	}

	songCategory.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Song category id cannot be parsed", err.Error()),
		)

		return
	}

	if err = songCategory.Get(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Song category not found", err.Error()),
		)

		return
	}

	c.JSON(200, utils.Success("Song category", songCategory))

}

func UpdateSongCategory(c *gin.Context) {

	var songCategory song.SongCategory
	var err error

	if err := c.ShouldBindJSON(&songCategory); err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Song category fields required", err.Error()),
		)

		return
	}

	if err = songCategory.Update(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Song category update failed", err.Error()),
		)

		return
	}

	c.JSON(200, utils.Success("Song category updated", songCategory))

}

func DeleteSongCategory(c *gin.Context) {

	var songCategory song.SongCategory
	var err error

	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Song category id required", "id must be required"),
		)
		return
	}

	songCategory.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error("Song category id cannot be parsed", err.Error()),
		)

		return
	}

	if err = songCategory.Delete(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Song category delete failed", err.Error()),
		)

		return
	}

	c.JSON(200, utils.Success("Song category deleted", songCategory))
}
