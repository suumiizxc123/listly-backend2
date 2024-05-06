package common_handler

import (
	"kcloudb1/internal/models/common"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateLanguage(c *gin.Context) {
	var language common.Language
	var err error

	if err := c.ShouldBindJSON(&language); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Language fields required", "Хэлний мэдээлэл дутуу байна"}, err),
		)
		return
	}

	if err = language.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Language creation failed", "Хэл хадгалахад алдаа үүслээ"}, err),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"Language created", "Хэл хадгаллаа"}, language))

}

func GetLanguageList(c *gin.Context) {

	var language common.Language

	languages, err := language.GetList()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Language list not found", "Хэлнийн жагсаалт олдсонгүй"}, err),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"Language list", "Хэлнийн жагсаалт"}, languages))
}

func UpdateLanguage(c *gin.Context) {
	var language common.Language
	var err error

	if err := c.ShouldBindJSON(&language); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Language fields required", "Хэлний мэдээлэл дутуу байна"}, err),
		)
		return
	}

	if err = language.Update(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Language update failed", "Хэл хадгалахад алдаа үүслээ"}, err),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"Language updated", "Хэл хадгаллаа"}, language))

}

func DeleteLanguage(c *gin.Context) {
	var language common.Language
	var err error

	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Language id required", "Хэлний id дутуу байна"}, err),
		)
		return
	}

	language.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Language id cannot be parsed", "Хэлний id буруу байна"}, err),
		)
		return
	}

	if err = language.Delete(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Language deletion failed", "Хэл устгахад алдаа үүслээ"}, err),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"Language deleted", "Хэл устгалаа"}, language))

}
