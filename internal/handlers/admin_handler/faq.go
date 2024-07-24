package admin_handler

import (
	"kcloudb1/internal/models/admin"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateFAQ(c *gin.Context) {
	var faq admin.FAQ

	if err := c.ShouldBindJSON(&faq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	faq.CreatedAt = time.Now()

	if err := faq.Create(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to create faq", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to create faq", "Амжилттай"}, faq))
}

func GetFAQList(c *gin.Context) {
	var faq admin.FAQ

	list, err := faq.GetList()
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get faq list", "Алдаа гарлаа"}, err))
		return
	}
	c.JSON(http.StatusOK, utils.Success([]string{"Success to get faq list", "Амжилттай"}, list))
}

func UpdateFAQ(c *gin.Context) {
	var faq admin.FAQ

	if err := c.ShouldBindJSON(&faq); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err))
		return
	}

	if err := faq.Update(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to update faq", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to update faq", "Амжилттай"}, faq))
}

func DeleteFAQ(c *gin.Context) {
	var faq admin.FAQ

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

	faq.ID = int64(idInt)

	if err := faq.Delete(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to delete faq", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to delete faq", "Амжилттай"}, nil))
}

func GetFAQ(c *gin.Context) {
	var faq admin.FAQ

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

	faq.ID = int64(idInt)

	if err := faq.Get(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get faq", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to get faq", "Амжилттай"}, faq))
}
