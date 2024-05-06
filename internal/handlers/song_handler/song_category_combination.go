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
			utils.Error([]string{"Song category combination fields required", "Дууны төрлийн холбоос мэдээлэл дутуу"}, err),
		)

		return
	}

	if err := songCategoryCombination.Create(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Song category combination creation failed", "Дууны төрлийн холбоос мэдээлэл нэмэхад алдаа гарлаа"}, err),
		)

		return
	}

	c.JSON(http.StatusCreated, utils.Success([]string{"Song category combination created", "Дууны төрлийн холбоос мэдээлэл нэмэгдлээ"}, songCategoryCombination))

}

func GetSongCategoryCombinationList(c *gin.Context) {

	var songCategoryCombination song.SongCategoryCombination

	songCategoryCombinations, err := songCategoryCombination.GetList()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Song category combination list retrieval failed", "Дууны төрлийн холбоос мэдээлэл жагсаалт авахад алдаа гарлаа"}, err),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"Song category combination list", "Дууны төрлийн холбоос мэдээлэл жагсаалт"}, songCategoryCombinations))
}

func GetSongCategoryCombination(c *gin.Context) {

	var songCategoryCombination song.SongCategoryCombination
	var err error
	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Song category combination id required", "Дууны төрлийн холбоос id мэдээлэл дутуу"}, "id must be required"),
		)

		return
	}

	songCategoryCombination.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Song category combination id cannot be parsed", "Дууны төрлийн холбоос id буруу байна"}, err.Error()),
		)

		return
	}

	if err = songCategoryCombination.Get(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Song category combination retrieval failed", "Дууны төрлийн холбоос мэдээлэл авахад алдаа гарлаа"}, err),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"Song category combination", "Дууны төрлийн холбоос мэдээлэл татлаа"}, songCategoryCombination))

}

func UpdateSongCategoryCombination(c *gin.Context) {

	var songCategoryCombination song.SongCategoryCombination

	if err := c.ShouldBindJSON(&songCategoryCombination); err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Song category combination fields required", "Дууны төрлийн холбоос мэдээлэл дутуу"}, err),
		)

		return
	}

	if err := songCategoryCombination.Update(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Song category combination update failed", "Дууны төрлийн холбоос мэдээлэл нэмэхад алдаа гарлаа"}, err),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"Song category combination updated", "Дууны төрлийн холбоос мэдээлэл нэмэгдлээ"}, songCategoryCombination))
}

func DeleteSongCategoryCombination(c *gin.Context) {

	var songCategoryCombination song.SongCategoryCombination
	var err error
	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Song category combination id required", "Дууны төрлийн холбоос id мэдээлэл дутуу"}, "id must be required"),
		)

		return
	}

	songCategoryCombination.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Song category combination id cannot be parsed", "Дууны төрлийн холбоос id буруу байна"}, err.Error()),
		)

		return
	}

	if err = songCategoryCombination.Delete(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Song category combination deletion failed", "Дууны төрлийн холбоос мэдээлэл устгахад алдаа гарлаа"}, err),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"Song category combination deleted", "Дууны төрлийн холбоос мэдээлэл устгагдлаа"}, songCategoryCombination))
}
