package order_handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kcloudb1/internal/config"
	"kcloudb1/internal/models/metal"
	"kcloudb1/internal/models/order"
	"kcloudb1/internal/models/payment"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

func CreateOrder(c *gin.Context) {
	var input order.CreateOrderInput
	var ord order.Order
	var ordp order.OrderPayment
	var met metal.MetalRate

	if err := c.ShouldBindJSON(&input); err != nil {
		resp := utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	clientIDStr := c.MustGet("clientID")

	clientID, err := strconv.ParseInt(clientIDStr.(string), 10, 64)

	if err != nil {
		resp := utils.Error([]string{"Failed to get clientID parse int64", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	// ord.Amount = input.Amount
	ord.Quantity = input.Quantity
	ord.MetalID = input.MetalID
	ord.ClientID = clientID

	tx := config.DB.Begin()

	if err := met.LastByMetalID(input.MetalID); err != nil {
		tx.Rollback()
	}

	ord.Price = met.Rate
	ord.Amount = ord.Price * ord.Quantity
	ord.Status = "pending"
	ord.CreatedAt = time.Now()

	if err := ord.Create(); err != nil {
		tx.Rollback()
	}

	ordp.OrderID = ord.ID
	ordp.InvoiceCode = "LISTLY_AGENT_INVOICE"
	newuid := uuid.NewString()
	ordp.SenderInvoiceNo = newuid
	ordp.InvoiceDescription = "GOLD PURCHASE"
	ordp.InvoiceReceiverCode = "terminal"
	ordp.SenderBranchCode = "SALBAR1"
	ordp.Amount = ord.Amount * 1.01
	ordp.CallbackURL = "http://oggbackend.999.mn:8080/api/v1/payment/" + newuid

	if err := ordp.Create(); err != nil {
		tx.Rollback()
	}

	res, err := sendInvoice(ordp)
	if err != nil {
		tx.Rollback()
	}

	tx.Commit()

	c.JSON(http.StatusOK, utils.Success([]string{"Success to create order", "Амжилттай"}, res))
}

func GetOrderList(c *gin.Context) {
	limit, _ := c.Get("limit")
	sort, _ := c.Get("sort")
	ordd, _ := c.Get("order")
	offset, _ := c.Get("offset")

	limitInt, _ := strconv.Atoi(limit.(string))

	offsetInt, _ := strconv.Atoi(offset.(string))
	clientIDStr := c.MustGet("clientID")

	clientID, err := strconv.ParseInt(clientIDStr.(string), 10, 64)

	if err != nil {
		resp := utils.Error([]string{"Failed to get clientID parse int64", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	var ord []order.OrderExtend

	if err := config.DB.Where("client_id = ?", clientID).Preload(clause.Associations).Limit(limitInt).Order(fmt.Sprintf("%s %s", sort.(string), ordd.(string))).Offset(offsetInt).Find(&ord).Error; err != nil {
		resp := utils.Error([]string{"Failed to get orders", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to get orders", "Амжилттай"}, ord))

}

func sendInvoice(orderp order.OrderPayment) (payment.QPayInvoiceResponse, error) {

	var ttk payment.QPayToken
	ttk.Last()

	var input payment.QPayInvoiceInput
	input.InvoiceCode = orderp.InvoiceCode
	input.SenderInvoiceNo = orderp.SenderInvoiceNo
	input.InvoiceReceiverCode = orderp.InvoiceReceiverCode
	input.InvoiceDescription = orderp.InvoiceDescription
	input.SenderBranchCode = orderp.SenderBranchCode
	input.Amount = orderp.Amount
	input.CallbackURL = orderp.CallbackURL

	// Create a map to represent the JSON payload

	// Convert the map to JSON
	jsonData, err := json.Marshal(input)
	if err != nil {
		return payment.QPayInvoiceResponse{}, err
	}

	fmt.Println("jsonData:", string(jsonData))

	// Create the request URL
	url := "https://merchant.qpay.mn/v2/invoice"

	// Create a new HTTP request with the POST method, URL, and request body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return payment.QPayInvoiceResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("Authorization", "Bearer "+ttk.AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return payment.QPayInvoiceResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return payment.QPayInvoiceResponse{}, err
	}

	fmt.Println("Response status:", resp.Status)
	var res payment.QPayInvoiceResponse
	err = json.Unmarshal(body, &res)

	if err != nil {
		fmt.Println("Error qpay:", err)
	}

	if err := config.DB.Model(&order.OrderPayment{}).
		Where("id = ?", orderp.ID).
		Update("invoice_id", res.InvoiceID).Error; err != nil {
		return payment.QPayInvoiceResponse{}, err
	}

	return res, nil

}
