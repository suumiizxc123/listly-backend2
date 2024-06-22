package order_handler

import (
	"kcloudb1/internal/models/metal"
	"kcloudb1/internal/models/order"
	"kcloudb1/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) {
	clientIDStr := c.MustGet("clientID")

	var bal order.Balance
	var balout order.BalanceResponse
	var met metal.MetalRate

	if err := bal.GetByClientAndMetalID(clientIDStr, 1); err != nil {
		resp := utils.Error([]string{"Failed to get balance", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if err := met.LastByMetalID(bal.MetalID); err != nil {
		resp := utils.Error([]string{"Failed to get metal", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	balout.ClientID = bal.ClientID
	balout.MetalID = bal.MetalID
	balout.Quantity = bal.Quantity
	balout.Balance = bal.Quantity * met.Rate
	balout.Changes = met.ChangePercent1D
	balout.CreatedAt = bal.CreatedAt

	resp := utils.Success([]string{"Success to get balance", "Амжилттай "}, balout)
	c.JSON(http.StatusOK, resp)
}

func GetBalanceHistory(c *gin.Context) {

	clientIDStr := c.MustGet("clientID")

	var balh order.BalanceHistory
	var resp utils.Response

	list, err := balh.GetByClientAndMetalID(clientIDStr, 1)

	if err != nil {
		resp = utils.Error([]string{"Failed to get balance history", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp = utils.Success([]string{"Success to get balance history", "Амжилттай"}, list)
	c.JSON(http.StatusOK, resp)

}
