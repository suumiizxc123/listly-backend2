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
			utils.Error("Language fields required", err),
		)
		return
	}

	if err = language.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Language creation failed", err),
		)
		return
	}

	c.JSON(200, utils.Success("Language created", language))

}

func GetLanguageList(c *gin.Context) {

	var language common.Language

	languages, err := language.GetList()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Language list not found", err),
		)
		return
	}

	c.JSON(200, utils.Success("Language list", languages))
}

func UpdateLanguage(c *gin.Context) {
	var language common.Language
	var err error

	if err := c.ShouldBindJSON(&language); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("Language fields required", err),
		)
		return
	}

	if err = language.Update(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Language update failed", err),
		)
		return
	}

	c.JSON(200, utils.Success("Language updated", language))

}

func DeleteLanguage(c *gin.Context) {
	var language common.Language
	var err error

	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("Language id required", err),
		)
		return
	}

	language.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("Language id cannot be parsed", err),
		)
		return
	}

	if err = language.Delete(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Language deletion failed", err),
		)
		return
	}

	c.JSON(200, utils.Success("Language deleted", language))

}
