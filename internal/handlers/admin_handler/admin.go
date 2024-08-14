package admin_handler

import (
	"fmt"
	"kcloudb1/internal/config"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"

	"kcloudb1/internal/models/admin"
	"kcloudb1/internal/models/client"
	"kcloudb1/internal/models/metal"
	"kcloudb1/internal/models/order"
	"kcloudb1/internal/models/saving"
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

func VerifyOrder(c *gin.Context) {
	orderID := c.Query("order_id")

	if orderID == "" {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get order id", "Алдаа гарлаа"}, nil))
		return
	}

	orderIDint, err := strconv.ParseInt(orderID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to convert order id", "Алдаа гарлаа"}, err))
		return
	}

	ord := order.Order{ID: orderIDint, AdminStatus: "success"}

	if err := ord.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to update order", "Алдаа гарлаа"}, err))
		return
	}

	if err := ord.Get(); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to get order", "Алдаа гарлаа"}, err))
		return
	}

	var bal order.Balance
	if ord.Type == "deposit" {
		// add balance

		if err := bal.GetByClientAndMetalID(ord.ClientID, 1); err != nil {
			c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to get balance", "Алдаа гарлаа"}, err.Error()))
			return
		}

		if err := config.DB.Model(&order.Balance{}).Where("id = ?", bal.ID).Update("quantity", bal.Quantity+ord.Quantity).Error; err != nil {
			c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to update balance", "Алдаа гарлаа"}, err.Error()))
			return
		}
	} else {
		// sub balance

		if err := bal.GetByClientAndMetalID(ord.ClientID, 1); err != nil {
			c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to get balance", "Алдаа гарлаа"}, err.Error()))
		}

		if bal.Quantity < ord.Quantity {
			c.JSON(http.StatusConflict, utils.Error([]string{"Out of balance", "Хүсэлтийн дүн хэтэрхий өндөр байна"}, "out of balance"))
			return
		}

		if err := config.DB.Model(&order.Balance{}).Where("id = ?", bal.ID).Update("quantity", bal.Quantity-ord.Quantity).Error; err != nil {
			c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to update balance", "Алдаа гарлаа"}, err.Error()))
			return
		}
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to update order", "Амжилттай"}, nil))
}

func CancelOrder(c *gin.Context) {
	orderID := c.Query("order_id")

	if orderID == "" {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get order id", "Алдаа гарлаа"}, nil))
		return
	}

	orderIDint, err := strconv.ParseInt(orderID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to convert order id", "Алдаа гарлаа"}, err))
		return
	}

	ord := order.Order{ID: orderIDint, AdminStatus: "cancel"}

	if err := ord.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to update order", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to update order", "Амжилттай"}, nil))
}

func CreateOrder(c *gin.Context) {
	var input struct {
		ClientID int64   `json:"client_id"`
		MetalID  int64   `json:"metal_id"`
		Quantity float32 `json:"quantity"`
		Amount   float32 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err))
		return
	}

	Price := input.Amount/input.Quantity - 50000.0

	ord := order.Order{
		ClientID:    input.ClientID,
		MetalID:     input.MetalID,
		Quantity:    input.Quantity,
		Amount:      input.Amount,
		Price:       Price,
		Type:        "deposit",
		CreatedAt:   time.Now(),
		Status:      "success",
		AdminStatus: "success",
	}

	if err := ord.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to create order", "Алдаа гарлаа"}, err))
		return
	}

	var bal order.Balance
	if err := bal.GetByClientAndMetalID(ord.ClientID, 1); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to get balance", "Алдаа гарлаа"}, err))
		return
	}

	if err := config.DB.Model(&order.Balance{}).Where("id = ?", bal.ID).Update("quantity", bal.Quantity+ord.Quantity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to update balance", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to create order", "Амжилттай"}, nil))
}

func VerifyWithDraw(c *gin.Context) {
	withdrawID := c.Query("withdraw_id")

	if withdrawID == "" {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get withdraw id", "Алдаа гарлаа"}, nil))
		return
	}

	withdrawIDint, err := strconv.ParseInt(withdrawID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to convert withdraw id", "Алдаа гарлаа"}, err))
		return
	}

	/// check balance and retrieve

	withdraw := order.WithDraw{ID: withdrawIDint, AdminStatus: "success"}

	if err := withdraw.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to update withdraw", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to update withdraw", "Амжилттай"}, nil))
}

func GetWithDrawList(c *gin.Context) {
	limit, _ := c.Get("limit")
	sort, _ := c.Get("sort")
	ord, _ := c.Get("order")
	offset, _ := c.Get("offset")

	limitInt, _ := strconv.Atoi(limit.(string))

	offsetInt, _ := strconv.Atoi(offset.(string))

	var resp utils.Response
	var wdr []order.WithDrawExtend

	if err := config.DB.Limit(limitInt).Order(fmt.Sprintf("%s %s", sort.(string), ord.(string))).Offset(offsetInt).Preload(clause.Associations).Find(&wdr).Error; err != nil {
		resp = utils.Error([]string{"Failed to get withdraws", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp = utils.Success([]string{"Success to get withdraws", "Амжилттай"}, wdr)
	c.JSON(http.StatusOK, resp)
}

func SendMessage(c *gin.Context) {
	var input struct {
		Message string `json:"message"`
		Phone   string `json:"phone"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err))
		return
	}

	if input.Message == "" || input.Phone == "" {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to send message", "Алдаа гарлаа"}, fmt.Errorf("message or phone not found")))
		return
	}

	err := utils.SendMessage(input.Phone, input.Message)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to send message", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to send message", "Амжилттай"}, nil))
}

func VerifySaving(c *gin.Context) {
	savingID := c.Query("saving_id")

	if savingID == "" {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to get saving id", "Алдаа гарлаа"}, nil))
		return
	}

	savingIDint, err := strconv.ParseInt(savingID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to convert saving id", "Алдаа гарлаа"}, err))
		return
	}

	/// check balance and retrieve

	saving := saving.SavingOrder{ID: savingIDint, AdminStatus: "success"}

	if err := saving.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to update saving", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to update saving", "Амжилттай"}, nil))
}
