package metal_handler

import (
	"kcloudb1/internal/models/metal"
	"kcloudb1/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLastMetalRate(c *gin.Context) {
	metal_id, ok := c.GetQuery("metal_id")
	var resp utils.Response
	var rate metal.MetalRate
	if !ok {
		resp = utils.Error([]string{"Failed to get metal", "Алдаа гарлаа"}, ok)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	err := rate.LastByMetalID(metal_id)
	if err != nil {
		resp = utils.Error([]string{"Failed to get rate", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp = utils.Success([]string{"Success to get rate", "Амжилттай"}, rate)
	c.JSON(http.StatusOK, resp)
}

func GetMetalRateByStartToEnd(c *gin.Context) {
	var resp utils.Response
	var rate metal.MetalRate

	metal_id, ok := c.GetQuery("metal_id")
	if !ok {
		resp = utils.Error([]string{"Failed to get metal", "Алдаа гарлаа"}, ok)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	start_date, ok := c.GetQuery("start_date")
	if !ok {
		resp = utils.Error([]string{"Failed to get start_date", "Алдаа гарлаа"}, ok)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	end_date, ok := c.GetQuery("end_date")
	if !ok {
		resp = utils.Error([]string{"Failed to get end_date", "Алдаа гарлаа"}, ok)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	order, ok := c.GetQuery("order")
	if !ok {
		order = "desc"
	}

	rates, err := rate.GetMetalRateByStartToEnd(metal_id, start_date, end_date, order)
	if err != nil {
		resp = utils.Error([]string{"Failed to get rates", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp = utils.Success([]string{"Success to get rates", "Амжилттай"}, rates)
	c.JSON(http.StatusOK, resp)
}

func GetMetalRateByKey(c *gin.Context) {
	key, ok := c.GetQuery("key")

	if !ok {
		resp := utils.Error([]string{"Failed to get key", "Алдаа гарлаа"}, ok)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	metal_id, ok := c.GetQuery("metal_id")
	if !ok {
		resp := utils.Error([]string{"Failed to get metal_id", "Алдаа гарлаа"}, ok)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	var order string
	order, ok = c.GetQuery("order")
	if !ok {
		order = "desc"
	}

	var rate metal.MetalRate
	var resp utils.Response

	rates, err := rate.GetMetalRateByKey(metal_id, key, order)
	if err != nil {
		resp = utils.Error([]string{"Failed to get rates", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp = utils.Success([]string{"Success to get rates", "Амжилттай"}, rates)
	c.JSON(http.StatusOK, resp)
}

func CreateMetalRate(c *gin.Context) {
	var metalRate metal.MetalRate
	var metalRateLast metal.MetalRate
	var resp utils.Response
	if err := c.ShouldBindJSON(&metalRate); err != nil {
		resp = utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	metalRateLasts, err := metalRateLast.GetMetalRateByKey(metalRate.MetalID, "last", "desc")

	if err != nil {
		resp = utils.Error([]string{"Failed to get last rate", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(metalRateLasts) > 0 {
		metalRateLast = metalRateLasts[0]
	}

	metalRate.ChangePercent1D = (metalRate.Rate - metalRateLast.Rate) / metalRateLast.Rate * 100

	if err := metalRate.Create(); err != nil {
		resp = utils.Error([]string{"Failed to create rate", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp = utils.Success([]string{"Success to create rate", "Амжилттай"}, nil)
	c.JSON(http.StatusOK, resp)
}
