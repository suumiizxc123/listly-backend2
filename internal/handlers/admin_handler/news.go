package admin_handler

import (
	"kcloudb1/internal/models/admin"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateNews(c *gin.Context) {
	var ns admin.News

	if err := c.ShouldBindJSON(&ns); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err))
		return
	}

	ns.CreatedAt = time.Now()

	if err := ns.Create(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to create news", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to create news", "Амжилттай"}, ns))
}

func UpdateNews(c *gin.Context) {
	var ns admin.News

	if err := c.ShouldBindJSON(&ns); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err))
		return
	}

	if err := ns.Update(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to update news", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to update news", "Амжилттай"}, ns))
}

func GetNewsList(c *gin.Context) {
	var ns admin.News

	news, err := ns.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get news", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to get news", "Амжилттай"}, news))
}

func DeleteNews(c *gin.Context) {
	var ns admin.News

	id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get id", "id дутуу байна"}, ok))
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to convert id", "id хөрвүүлэлтэнд алдаа гарлаа"}, err))
		return
	}

	ns.ID = int64(idInt)

	if err := ns.Delete(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to delete news", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to delete news", "Амжилттай"}, nil))
}

func GetNews(c *gin.Context) {
	var ns admin.News

	id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get id", "id дутуу байна"}, ok))
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to convert id", "id хөрвүүлэлтэнд алдаа гарлаа"}, err))
		return
	}

	ns.ID = int64(idInt)

	if err := ns.Get(); err != nil {

		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get news", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to get news", "Амжилттай"}, ns))
}
