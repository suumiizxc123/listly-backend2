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
			utils.Error("Service log fields required", err.Error()),
		)
		return
	}

	if err = log.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Service log creation failed", err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success("Service log created", log))
}

func GetServiceLogList(c *gin.Context) {
	var log user.ServiceLog
	var err error

	logs, err := log.GetList()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Service log list not found", err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success("Service log list", logs))
}

func UpdateServiceLog(c *gin.Context) {
	var log user.ServiceLog
	var err error

	if err = c.ShouldBindJSON(&log); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("Service log fields required", err.Error()),
		)
		return
	}

	if err = log.Update(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Service log update failed", err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success("Service log updated", log))
}

func DeleteServiceLog(c *gin.Context) {
	var log user.ServiceLog
	var err error

	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("Service log id required", "id must be required"),
		)
		return
	}

	log.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error("Service log id cannot be parsed", err.Error()),
		)
		return
	}

	if err = log.Delete(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error("Service log delete failed", err.Error()),
		)
		return
	}

	c.JSON(200, utils.Success("Service log deleted", struct{}{}))
}
