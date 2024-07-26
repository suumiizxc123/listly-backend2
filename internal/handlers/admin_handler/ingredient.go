package admin_handler

import (
	"kcloudb1/internal/models/admin"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateIngredient(c *gin.Context) {
	var ingr admin.Ingredient
	err := c.ShouldBindJSON(&ingr)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to create ingredient", "Алдаа гарлаа"}, err))
		return
	}
	err = ingr.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to create ingredient", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to create ingredient", "Амжилттай"}, ingr))
}

func UpdateIngredient(c *gin.Context) {
	var ingr admin.Ingredient
	err := c.ShouldBindJSON(&ingr)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to update ingredient", "Алдаа гарлаа"}, err))
		return
	}
	err = ingr.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to update ingredient", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to update ingredient", "Амжилттай"}, ingr))
}

func GetIngredientList(c *gin.Context) {
	var ingr admin.Ingredient
	ingrs, err := ingr.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get ingredient list", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to get ingredient list", "Амжилттай"}, ingrs))
}

func DeleteIngredient(c *gin.Context) {
	var ingr admin.Ingredient

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

	ingr.ID = int64(idInt)

	if err := ingr.Delete(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to delete ingredient", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to delete ingredient", "Амжилттай"}, nil))
}

func GetIngredient(c *gin.Context) {
	var ingr admin.Ingredient

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

	ingr.ID = int64(idInt)

	if err := ingr.Get(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get ingredient", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to get ingredient", "Амжилттай"}, ingr))
}
