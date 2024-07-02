package order_handler

import (
	"fmt"
	"kcloudb1/internal/models/metal"
	"kcloudb1/internal/models/order"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateWithdraw(c *gin.Context) {
	var input order.CreateWithDrawInput
	var wdr order.WithDraw
	var met metal.MetalRate
	var bal order.Balance
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to bind json", "Мэдээлэл дутуу байна"}, err))
		return
	}

	clientIDStr := c.MustGet("clientID")

	clientID, err := strconv.ParseInt(clientIDStr.(string), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get clientID parse int64", "Хэрэглэгчийн ID дуудаагүй байна"}, err))
		return
	}

	wdr.ClientID = clientID
	wdr.Quantity = input.Quantity
	wdr.MetalID = input.MetalID
	wdr.Type = "withdraw"

	if err := met.LastByMetalID(input.MetalID); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to get metal", "Үнийн мэдээлэл дээр алдаа үүслээ"}, err.Error()))
		return
	}

	wdr.Price = met.Rate
	wdr.Amount = wdr.Price * wdr.Quantity

	if err := bal.GetByClientAndMetalID(clientID, 1); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to get balance", "Хэрэглэгчийн баланс дээр алдаа үүслээ"}, err.Error()))
		return
	}

	if bal.Quantity < wdr.Quantity {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to withdraw", "Хэрэглэгчийн үлдэгдэл хүрэлцэхгүй байна"}, fmt.Errorf("balance is not enough")))
		return
	}

	wdr.AdminStatus = "pending"
	wdr.Status = "success"

	if err := wdr.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to create withdraw", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to create withdraw", "Амжилттай"}, wdr))
}

func GetWithdraw(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get id", "id дутуу байна"}, nil))
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to convert id", "id хөрвүүлэлтэнд алдаа гарлаа"}, err))
		return
	}

	wdr := order.WithDrawExtend{ID: id}

	if err := wdr.Get(); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to get withdraw", "Мэдээллийг татахад алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to get withdraw", "Амжилттай"}, wdr))
}

func GetWithDrawList(c *gin.Context) {

	clientIDStr := c.MustGet("clientID")
	clientID, err := strconv.ParseInt(clientIDStr.(string), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get clientID parse int64", "Хэрэглэгчийн ID дуудаагүй байна"}, err))
		return
	}

	wdr := order.WithDrawExtend{}
	wdrs, err := wdr.GetListByClientID(clientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to get withdraw list", "Мэдээллийг татахад алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to get withdraw list", "Амжилттай"}, wdrs))
}
