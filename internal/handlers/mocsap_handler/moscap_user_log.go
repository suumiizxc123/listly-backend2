package moscap_handler

import (
	"kcloudb1/internal/models/moscap"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateMoscapUserLog(c *gin.Context) {
	var org moscap.MosCapUserLog
	var err error

	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Org fields required", " мэдээлэл дутуу байна"}, err),
		)
		return
	}

	org.CreatedAt = time.Now()

	if err = org.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Org creation failed", " бүртгэхэд алдаа гарлаа"}, err),
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		utils.Success([]string{"Org created", " бүртгэгдлээ"}, org),
	)
}

func UpdateMosCapUserLog(c *gin.Context) {
	var org moscap.MosCapUserLog
	var err error

	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Org fields required", " мэдээлэл дутуу байна"}, err),
		)
		return
	}

	if err = org.Update(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Org update failed", " бүртгэхэд алдаа гарлаа"}, err),
		)
		return
	}

	c.JSON(
		200,
		utils.Success([]string{"Org updated", " бүртгэгдлээ"}, org),
	)
}

func DeleteMosCapUserLog(c *gin.Context) {
	var org moscap.MosCapUserLog
	var err error

	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Org id required", " id дутуу байна"}, err),
		)
		return
	}

	org.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Org id cannot be parsed", " id буруу байна"}, err),
		)
		return
	}

	if err = org.Delete(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Org deletion failed", " устгахад алдаа гарлаа"}, err),
		)
		return
	}

	c.JSON(
		200,
		utils.Success([]string{"Org deleted", " устгагдлаа"}, struct{}{}),
	)

}

func GetMosCapUserLog(c *gin.Context) {
	var org moscap.MosCapUserLog
	var err error

	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Org id required", " id дутуу байна"}, err),
		)
		return
	}

	org.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Org id cannot be parsed", " id буруу байна"}, err),
		)
		return
	}

	if err = org.Get(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Org not found", " буруу байна"}, err),
		)
		return
	}

	c.JSON(
		200,
		utils.Success([]string{"Org found", " бүртгэгдлээ"}, org),
	)

}

func GetMosCapUserLogList(c *gin.Context) {
	var org moscap.MosCapUserLog
	var err error

	orgs, err := org.GetList()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Org list not found", " жагсаалт буруу байна"}, err),
		)
		return
	}

	c.JSON(
		200,
		utils.Success([]string{"Org list", " жагсаалт"}, orgs),
	)
}
