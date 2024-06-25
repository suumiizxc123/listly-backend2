package order_handler

import (
	"bytes"
	"encoding/json"
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
	ord.Amount = input.Amount
	ord.MetalID = input.MetalID
	ord.ClientID = clientID

	tx := config.DB.Begin()

	if err := met.LastByMetalID(input.MetalID); err != nil {
		tx.Rollback()
	}

	ord.Price = met.Rate
	ord.Quantity = input.Amount / met.Rate
	ord.Status = "pending"
	ord.CreatedAt = time.Now()

	if err := ord.Create(); err != nil {
		tx.Rollback()
	}

	ordp.OrderID = ord.ID
	ordp.InvoiceCode = "TEST_INVOICE"
	newuid := uuid.NewString()
	ordp.SenderInvoiceNo = newuid
	ordp.InvoiceDescription = "test"
	ordp.InvoiceReceiverCode = "terminal"
	ordp.SenderBranchCode = "SALBAR1"
	ordp.Amount = ord.Amount
	ordp.CallbackURL = "https://3fc9-202-21-120-182.ngrok-free.app/payments/" + newuid

	if err := ordp.Create(); err != nil {
		tx.Rollback()
	}

	if err := sendInvoice(ordp); err != nil {
		tx.Rollback()
	}

	tx.Commit()

	c.JSON(http.StatusOK, utils.Success([]string{"Success to create order", "Амжилттай"}, nil))

}

func sendInvoice(orderp order.OrderPayment) error {
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
		return err
	}

	// Create the request URL
	url := "https://merchant-sandbox.qpay.mn/v2/invoice"

	// Create a new HTTP request with the POST method, URL, and request body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")
	// Set the x-api-key header
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRfaWQiOiI5OWNmYjFiMi0zMmNjLTRhNDMtOWNjZi02NWNhM2JjOTg0YWYiLCJzZXNzaW9uX2lkIjoiQVhXQnNtUXJmb2hObE0xMW9Vd1c5M3NYc2tleUNnbG8iLCJpYXQiOjE3MTgwMzIyNDMsImV4cCI6MzQzNjE1MDg4Nn0.F-bmsKGihb7RHcj9Pbf22LjeH86mYaYvNuIp3-KAr0A")

	// Send the request using the default HTTP client
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil

}
