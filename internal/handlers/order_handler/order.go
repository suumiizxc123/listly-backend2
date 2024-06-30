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
	timeStart := time.Now()
	var input order.CreateOrderInput
	var ord order.Order
	// var ordp order.OrderPayment
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

	ord.Quantity = input.Quantity
	ord.MetalID = input.MetalID
	ord.ClientID = clientID
	if err := met.LastByMetalID(input.MetalID); err != nil {
		resp := utils.Error([]string{"Failed to get metal", "Алдаа гарлаа"}, err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	ord.Price = met.Rate
	ord.Amount = ord.Price*ord.Quantity + 50000.0
	ord.Status = "pending"
	ord.CreatedAt = time.Now()

	fmt.Println("step 1 : ", time.Now().Sub(timeStart))
	if err := ord.Create(); err != nil {
		tx.Rollback()
	}
	newuid := uuid.NewString()
	ordp := order.OrderPayment{
		OrderID:             ord.ID,
		InvoiceCode:         "LISTLY_AGENT_INVOICE",
		SenderInvoiceNo:     newuid,
		InvoiceDescription:  "GOLD PURCHASE",
		InvoiceReceiverCode: "terminal",
		SenderBranchCode:    "SALBAR1",
		Amount:              ord.Amount / 0.99,
		CallbackURL:         fmt.Sprintf("http://oggbackend.999.mn:8080/api/v1/payment/%s", newuid),
	}

	if err := ordp.Create(); err != nil {
		tx.Rollback()
	}
	fmt.Println("step 2 : ", time.Now().Sub(timeStart))
	res, err := sendInvoice(ordp)
	if err != nil {
		tx.Rollback()
	}

	// Prepare response
	resm := map[string]interface{}{
		"invoice_id":    res.InvoiceID,
		"qr_text":       res.QRText,
		"qr_image":      res.QRImage,
		"qPay_shortUrl": res.QPayShortUrl,
		"urls":          res.Urls,
		"amount":        ord.Amount,
		"quantity":      ord.Quantity,
	}
	fmt.Println("step 3 : ", time.Now().Sub(timeStart))
	tx.Commit()

	c.JSON(http.StatusOK, utils.Success([]string{"Success to create order", "Амжилттай"}, resm))
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
func sendInvoice(ordp order.OrderPayment) (payment.QPayInvoiceResponse, error) {
	var ttk payment.QPayToken
	ttk.Last()

	input := payment.QPayInvoiceInput{
		InvoiceCode:         ordp.InvoiceCode,
		SenderInvoiceNo:     ordp.SenderInvoiceNo,
		InvoiceReceiverCode: ordp.InvoiceReceiverCode,
		InvoiceDescription:  ordp.InvoiceDescription,
		SenderBranchCode:    ordp.SenderBranchCode,
		Amount:              ordp.Amount,
		CallbackURL:         ordp.CallbackURL,
	}

	// Marshal JSON input
	jsonData, err := json.Marshal(input)
	if err != nil {
		return payment.QPayInvoiceResponse{}, err
	}

	// Prepare HTTP request
	url := "https://merchant.qpay.mn/v2/invoice"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return payment.QPayInvoiceResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ttk.AccessToken)

	// Execute HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return payment.QPayInvoiceResponse{}, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return payment.QPayInvoiceResponse{}, err
	}

	// Parse JSON response
	var res payment.QPayInvoiceResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return payment.QPayInvoiceResponse{}, err
	}

	// Update order payment with invoice ID
	if err := config.DB.Model(&order.OrderPayment{}).
		Where("id = ?", ordp.ID).
		Update("invoice_id", res.InvoiceID).Error; err != nil {
		return payment.QPayInvoiceResponse{}, err
	}

	return res, nil
}
