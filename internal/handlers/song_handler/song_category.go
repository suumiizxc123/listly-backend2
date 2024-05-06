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
			utils.Error([]string{"Song category fields required", "Дууны төрлийн мэдээлэл дутуу байна"}, err),
		)

		return
	}

	if err = songCategory.Create(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Song category creation failed", "Дууны төрлийн мэдээлэл хадгалахад алдаа үүслээ"}, err),
		)

		return
	}

	c.JSON(http.StatusCreated, utils.Success([]string{"Song category created", "Дууны төрлийн мэдээлэл хадгаллаа"}, songCategory))
}

func GetSongCategoryList(c *gin.Context) {

	var songCategory song.SongCategory

	songCategories, err := songCategory.GetList()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Song category list not found", "Дууны төрлийн мэдээлэл татахад үүслээ"}, err.Error()),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"Song category list", "Дууны төрлийн мэдээлэл татлаа"}, songCategories))
}

func GetSongCategory(c *gin.Context) {

	var songCategory song.SongCategorySong
	var err error
	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Song category id required", "Дууны төрлийн холбоос id мэдээлэл дутуу"}, "id must be required"),
		)

		return

	}

	songCategory.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Song category id cannot be parsed", "Дууны төрлийн холбоос id буруу байна"}, err.Error()),
		)

		return
	}

	if err = songCategory.Get(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Song category not found", "Дууны төрлийн мэдээлэл буруу байна"}, err.Error()),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"Song category", "Дууны төрлийн мэдээлэл амжилттай татлаа"}, songCategory))

}

func UpdateSongCategory(c *gin.Context) {

	var songCategory song.SongCategory
	var err error

	if err := c.ShouldBindJSON(&songCategory); err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Song category fields required", "Дууны төрлийн мэдээлэл дутуу байна"}, err.Error()),
		)

		return
	}

	if err = songCategory.Update(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Song category update failed", "Дууны төрлийн мэдээлэл хадгалахад алдаа үүслээ"}, err.Error()),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"Song category updated", "Дууны төрлийн мэдээлэл хадгаллаа"}, songCategory))

}

func DeleteSongCategory(c *gin.Context) {

	var songCategory song.SongCategory
	var err error

	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Song category id required", "Дууны төрлийн холбоос id мэдээлэл дутуу"}, "id must be required"),
		)
		return
	}

	songCategory.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Song category id cannot be parsed", "Дууны төрлийн холбоос id буруу байна"}, err.Error()),
		)

		return
	}

	if err = songCategory.Delete(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Song category delete failed", "Дууны төрлийн мэдээлэл хадгалахад алдаа үүслээ"}, err.Error()),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"Song category deleted", "Дууны төрлийн мэдээлэл хадгаллаа"}, songCategory))
}
