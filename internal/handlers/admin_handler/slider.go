package admin_handler

import (
	"kcloudb1/internal/models/admin"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateSlider(c *gin.Context) {
	var sl admin.Slider

	if err := c.ShouldBindJSON(&sl); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err))
		return
	}

	sl.CreatedAt = time.Now()

	if err := sl.Create(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to create slider", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to create slider", "Амжилттай"}, sl))
}

func GetSliderList(c *gin.Context) {
	var sl admin.Slider

	list, err := sl.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get slider list", "Алдаа гарлаа"}, err))
		return
	}
	c.JSON(http.StatusOK, utils.Success([]string{"Success to get slider list", "Амжилттай"}, list))
}

func GetSlider(c *gin.Context) {
	var sl admin.Slider

	id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get id", "id дутуу байна"}, ok))
		return
	}

	slIDint, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to convert id", "id хөрвүүлэлтэнд алдаа гарлаа"}, err))
		return
	}

	sl.ID = int64(slIDint)

	if err := sl.Get(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get slider", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to get slider", "Амжилттай"}, sl))
}

func UpdateSlider(c *gin.Context) {
	var sl admin.Slider

	if err := c.ShouldBindJSON(&sl); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err))
		return
	}

	if err := sl.Update(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to update slider", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to update slider", "Амжилттай"}, sl))
}

func DeleteSlider(c *gin.Context) {
	var sl admin.Slider

	id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get id", "id дутуу байна"}, ok))
		return
	}

	slIDint, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to convert id", "id хөрвүүлэлтэнд алдаа гарлаа"}, err))
		return
	}

	sl.ID = int64(slIDint)

	if err := sl.Delete(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to delete slider", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to delete slider", "Амжилттай"}, nil))
}
