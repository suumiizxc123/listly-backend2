package admin_handler

import (
	"fmt"
	"kcloudb1/internal/config"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"

	"kcloudb1/internal/models/admin"
	"kcloudb1/internal/models/client"
	"kcloudb1/internal/models/metal"
	"kcloudb1/internal/models/order"
)

func LoginByPassword(c *gin.Context) {
	var data struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	var adm admin.Admin
	var admout admin.AdminOutput
	var resp utils.Response
	if err := c.ShouldBindJSON(&data); err != nil {
		resp = utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	adm.Name = data.Name
	adm.Password = data.Password

	if err := adm.GetByName(data.Name); err != nil {
		resp = utils.Error([]string{"Failed to get admin by name", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if adm.IsActive == 0 {
		resp = utils.Error([]string{"Admin not active", "Алдаа гарлаа"}, fmt.Errorf("%v", "Admin not active"))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	token := uuid.NewString()

	result := config.RS.Set(token, adm.ID, 0)

	if result.Err() != nil {
		resp = utils.Error([]string{"Failed to set token", "Алдаа гарлаа"}, result.Err())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	admout.Token = token
	admout.ID = adm.ID
	admout.IsActive = adm.IsActive
	admout.Name = adm.Name

	resp = utils.Success([]string{"Success to login", "Амжилттай"}, admout)
	c.JSON(http.StatusOK, resp)
}

func GetClientList(c *gin.Context) {
	limit, _ := c.Get("limit")
	sort, _ := c.Get("sort")
	order, _ := c.Get("order")
	offset, _ := c.Get("offset")

	limitInt, _ := strconv.Atoi(limit.(string))

	offsetInt, _ := strconv.Atoi(offset.(string))
	var resp utils.Response
	var usrs []client.ClientOutput

	if err := config.DB.Limit(limitInt).Order(fmt.Sprintf("%s %s", sort, order)).Offset(offsetInt).Preload(clause.Associations).Find(&usrs).Error; err != nil {
		resp = utils.Error([]string{"Failed to get users", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp = utils.Success([]string{"Success to get users", "Амжилттай"}, usrs)
	c.JSON(http.StatusOK, resp)
}

func UpdateClient(c *gin.Context) {
	var cl client.Client

	if err := c.ShouldBindJSON(&cl); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err))
		return
	}

	if err := cl.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to update client", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to update client", "Амжилттай"}, nil))
}

func DeleteClient(c *gin.Context) {
	clientID := c.Query("client_id")

	if clientID == "" {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get client id", "Алдаа гарлаа"}, nil))
		return
	}

	clientIDint, err := strconv.ParseInt(clientID, 10, 64)
	if err != nil {

		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to convert client id", "Алдаа гарлаа"}, err))
		return
	}

	cl := client.Client{ID: clientIDint}

	if err := cl.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to delete client", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to delete client", "Амжилттай"}, nil))
}

func GetBalanceByClientID(c *gin.Context) {
	clientID := c.Query("client_id")

	if clientID == "" {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get client id", "Алдаа гарлаа"}, nil))
		return
	}

	var bal order.Balance
	var balout order.BalanceResponse
	var met metal.MetalRate

	if err := bal.GetByClientAndMetalID(clientID, 1); err != nil {
		resp := utils.Error([]string{"Failed to get balance", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if err := met.LastByMetalID(bal.MetalID); err != nil {
		resp := utils.Error([]string{"Failed to get metal", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	balout.ID = bal.ID
	balout.ClientID = bal.ClientID
	balout.MetalID = bal.MetalID
	balout.Quantity = bal.Quantity
	balout.Balance = bal.Quantity * met.Rate
	balout.Changes = met.ChangePercent1D
	balout.CreatedAt = bal.CreatedAt

	resp := utils.Success([]string{"Success to get balance", "Амжилттай "}, balout)
	c.JSON(http.StatusOK, resp)
}

func GetOrderList(c *gin.Context) {
	limit, _ := c.Get("limit")
	sort, _ := c.Get("sort")
	ord, _ := c.Get("order")
	offset, _ := c.Get("offset")

	limitInt, _ := strconv.Atoi(limit.(string))

	offsetInt, _ := strconv.Atoi(offset.(string))

	var resp utils.Response
	var ords []order.OrderExtend

	if err := config.DB.Limit(limitInt).Order(fmt.Sprintf("%s %s", sort.(string), ord.(string))).Offset(offsetInt).Preload(clause.Associations).Find(&ords).Error; err != nil {
		resp = utils.Error([]string{"Failed to get orders", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp = utils.Success([]string{"Success to get orders", "Амжилттай"}, ords)
	c.JSON(http.StatusOK, resp)
}
