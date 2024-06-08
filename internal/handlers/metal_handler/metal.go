package metal_handler

import (
	"kcloudb1/internal/models/metal"
	"kcloudb1/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMetal(c *gin.Context) {
	var metall metal.Metal
	var resp utils.Response
	if err := c.ShouldBindJSON(&metall); err != nil {
		resp = utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if err := metall.Create(); err != nil {
		resp = utils.Error([]string{"Failed to create metal", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp = utils.Success([]string{"Success to create metal", "Амжилттай"}, nil)
	c.JSON(http.StatusOK, resp)
}

func GetAllMetals(c *gin.Context) {
	var metall metal.Metal
	var resp utils.Response
	metals, err := metall.GetAll()
	if err != nil {
		resp = utils.Error([]string{"Failed to get metals", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp = utils.Success([]string{"Success to get metals", "Амжилттай"}, metals)
	c.JSON(http.StatusOK, resp)
}

func UpdateMetal(c *gin.Context) {
	var metall metal.Metal
	var resp utils.Response
	if err := c.ShouldBindJSON(&metall); err != nil {
		resp = utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if err := metall.Update(); err != nil {
		resp = utils.Error([]string{"Failed to update metal", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp = utils.Success([]string{"Success to update metal", "Амжилттай"}, nil)
	c.JSON(http.StatusOK, resp)
}

func DeleteMetal(c *gin.Context) {
	var metall metal.Metal
	var resp utils.Response
	id, ok := c.GetQuery("id")
	if !ok {
		resp = utils.Error([]string{"Failed to get id", "Алдаа гарлаа"}, ok)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if err := metall.Delete(id); err != nil {
		resp = utils.Error([]string{"Failed to delete metal", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp = utils.Success([]string{"Success to delete metal", "Амжилттай"}, nil)
	c.JSON(http.StatusOK, resp)
}

func GetMetal(c *gin.Context) {
	var metall metal.Metal
	var resp utils.Response
	id, ok := c.GetQuery("id")
	if !ok {
		resp = utils.Error([]string{"Failed to get id", "Алдаа гарлаа"}, ok)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if err := metall.Get(id); err != nil {
		resp = utils.Error([]string{"Failed to get metal", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp = utils.Success([]string{"Success to get metal", "Амжилттай"}, metall)
	c.JSON(http.StatusOK, resp)
}
