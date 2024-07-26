package admin_handler

import (
	"kcloudb1/internal/models/admin"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var pr admin.Product

	if err := c.ShouldBindJSON(&pr); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err))
		return
	}

	pr.CreatedAt = time.Now()

	if err := pr.Create(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to create product", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to create product", "Амжилттай"}, pr))
}

func UpdateProduct(c *gin.Context) {
	var pr admin.Product

	if err := c.ShouldBindJSON(&pr); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err))
		return
	}

	if err := pr.Update(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to update product", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to update product", "Амжилттай"}, pr))
}

func GetProductList(c *gin.Context) {
	var pr admin.ProductExtend

	prs, err := pr.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get product list", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to get product list", "Амжилттай"}, prs))
}

func GetProduct(c *gin.Context) {
	var pr admin.ProductExtend

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

	pr.ID = int64(idInt)

	if err := pr.Get(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get product", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to get product", "Амжилттай"}, pr))
}

func DeleteProduct(c *gin.Context) {
	var pr admin.Product

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

	pr.ID = int64(idInt)

	if err := pr.Delete(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to delete product", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to delete product", "Амжилттай"}, nil))
}

func AddProductImage(c *gin.Context) {
	var pr admin.Product
	var pi admin.ProductImage
	var input struct {
		ProductID int64  `json:"product_id"`
		Image     string `json:"image"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err))
		return
	}

	pi.ProductID = input.ProductID
	pi.Image = input.Image
	pi.CreatedAt = time.Now()

	pr.ID = input.ProductID

	if err := pr.Get(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get product", "Алдаа гарлаа"}, err))
		return
	}

	if err := pi.Create(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to create product image", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to create product image", "Амжилттай"}, pi))
}

func RemoveProductImage(c *gin.Context) {
	var pr admin.Product
	var pi admin.ProductImage
	var input struct {
		ProductID      int64 `json:"product_id"`
		ProductImageID int64 `json:"product_image_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err))
		return
	}

	pi.ID = input.ProductImageID
	pi.ProductID = input.ProductID

	pr.ID = input.ProductID

	if err := pr.Get(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get product", "Алдаа гарлаа"}, err))
		return
	}

	if err := pi.Delete(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to delete product image", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to delete product image", "Амжилттай"}, nil))
}

func AddProductIngredient(c *gin.Context) {
	var pr admin.Product
	var pi admin.ProductIngredient
	var input struct {
		ProductID    int64  `json:"product_id"`
		IngredientID int64  `json:"ingredient_id"`
		Description  string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err))
		return
	}

	pi.ProductID = input.ProductID
	pi.IngredientID = input.IngredientID
	pi.Description = input.Description
	pi.CreatedAt = time.Now()

	pr.ID = input.ProductID

	if err := pr.Get(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get product", "Алдаа гарлаа"}, err))
		return
	}

	if err := pi.Create(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to create product ingredient", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to create product ingredient", "Амжилттай"}, pi))
}

func RemoveProductIngredient(c *gin.Context) {
	var pr admin.Product
	var pi admin.ProductIngredient
	var input struct {
		ProductID           int64 `json:"product_id"`
		ProductIngredientID int64 `json:"product_ingredient_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err))
		return
	}

	pi.ID = input.ProductIngredientID
	pi.ProductID = input.ProductID

	pr.ID = input.ProductID

	if err := pr.Get(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get product", "Алдаа гарлаа"}, err))
		return
	}

	if err := pi.Delete(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to delete product ingredient", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to delete product ingredient", "Амжилттай"}, nil))
}
