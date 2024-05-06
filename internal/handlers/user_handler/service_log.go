package user_handler

import (
	"kcloudb1/internal/models/user"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateServiceLog(c *gin.Context) {
	var log user.ServiceLog
	var err error

	if err = c.ShouldBindJSON(&log); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Service log fields required", "Системийн лог мэдээлэл дутуу байна"}, err.Error()),
		)
		return
	}

	if err = log.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Service log creation failed", "Системийн лог мэдээлэл хадгалахад алдаа үүслээ"}, err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"Service log created", "Системийн лог мэдээлэл хадгаллаа"}, log))
}

func GetServiceLogList(c *gin.Context) {
	var log user.ServiceLog
	var err error

	logs, err := log.GetList()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Service log list not found", "Системийн лог лист байхгүй байна"}, err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"Service log list", "Системийн лог лист татлаа"}, logs))
}

func UpdateServiceLog(c *gin.Context) {
	var log user.ServiceLog
	var err error

	if err = c.ShouldBindJSON(&log); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Service log fields required", "Системийн лог мэдээлэл дутуу байна"}, err.Error()),
		)
		return
	}

	if err = log.Update(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Service log update failed", "Системийн лог мэдээлэл хадгалахад алдаа үүслээ"}, err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"Service log updated", "Системийн лог мэдээлэл хадгаллаа"}, log))
}

func DeleteServiceLog(c *gin.Context) {
	var log user.ServiceLog
	var err error

	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Service log id required", "Системийн лог мэдээлэл дутуу байна"}, "id must be required"),
		)
		return
	}

	log.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Service log id cannot be parsed", "Системийн лог мэдээлэл буруу байна"}, err.Error()),
		)
		return
	}

	if err = log.Delete(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Service log delete failed", "Системийн лог мэдээлэл устгахад алдаа үүслээ"}, err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"Service log deleted", "Системийн лог мэдээлэл устгалаа"}, struct{}{}))
}
