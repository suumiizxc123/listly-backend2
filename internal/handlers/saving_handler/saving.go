package saving_handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kcloudb1/internal/config"
	"kcloudb1/internal/models/client"
	"kcloudb1/internal/models/metal"
	"kcloudb1/internal/models/payment"
	"kcloudb1/internal/models/saving"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateSavingOrder(c *gin.Context) {
	var input saving.CreateSavingOrderInput
	var sa saving.SavingOrder
	var met metal.MetalRate
	var cl client.Client

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

	sa.ClientID = clientID

	sa.Term = input.Term
	sa.MetalID = input.MetalID
	sa.Status = "pending"

	cl.ID = clientID

	if err := cl.Get(); err != nil {
		resp := utils.Error([]string{"Failed to get client", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

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

	sa.Price = met.Rate
	sa.Quantity = input.Quantity
	sa.Amount = sa.Price * sa.Quantity
	sa.Status = "pending"
	sa.AdminStatus = "pending"
	sa.CreatedAt = time.Now()
	sa.Type = "deposit"
	sa.StartDate = time.Now()
	sa.EndDate = sa.StartDate.AddDate(0, 0, int(sa.Term))
	sa.TermGrowth = getRatio(sa.Term)
	sa.ExpectedAmount = sa.Amount * (1 + sa.TermGrowth)

	txncount := sa.GetClientTotalCount(clientID)
	sa.TransactionText = fmt.Sprintf("%v-%v", cl.Phone, txncount+1)

	if err := sa.Create(); err != nil {
		tx.Rollback()
		resp := utils.Error([]string{"Failed to create saving order", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	newuid := uuid.NewString()

	sadp := saving.SavingOrderPayment{
		SavingOrderID:       sa.ID,
		InvoiceCode:         "LISTLY_AGENT_INVOICE",
		SenderInvoiceNo:     newuid,
		InvoiceDescription:  "GOLD SAVING",
		InvoiceReceiverCode: "terminal",
		SenderBranchCode:    "SALBAR1",
		Amount:              sa.Amount,
		CallbackURL:         fmt.Sprintf("http://oggbackend.999.mn:8080/api/v1/payment/saving/%s", newuid),
	}

	if err := sadp.Create(); err != nil {
		tx.Rollback()
		resp := utils.Error([]string{"Failed to create saving order payment", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	res, err := sendInvoice(sadp)
	if err != nil {
		tx.Rollback()
		resp := utils.Error([]string{"Failed to send invoice", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resm := map[string]interface{}{
		"order_id":      sa.ID,
		"invoice_id":    res.InvoiceID,
		"qr_text":       res.QRText,
		"qr_image":      res.QRImage,
		"qPay_shortUrl": res.QPayShortUrl,
		"urls":          res.Urls,
		"amount":        sa.Amount,
		"quantity":      sa.Quantity,
		"transaction":   sa.TransactionText,
	}

	tx.Commit()

	c.JSON(http.StatusOK, utils.Success([]string{"Success to create saving order", "Амжилттай"}, resm))

}

func GetSavingOrder(c *gin.Context) {

	clientIDStr := c.MustGet("clientID")

	clientID, err := strconv.ParseInt(clientIDStr.(string), 10, 64)

	if err != nil {
		resp := utils.Error([]string{"Failed to get clientID parse int64", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	var sa saving.SavingOrder
	sa.ClientID = clientID

	sas, err := sa.GetByClientID(clientID)

	if err != nil {
		resp := utils.Error([]string{"Failed to get saving order", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	var output saving.SavingOutput

	output.SavingOrders = sas

	output.TotalAmount = 0
	output.TotalInvest = 0
	output.TotalProfit = 0

	for _, v := range sas {
		output.TotalAmount += v.Amount
		output.TotalInvest += v.Quantity
		output.TotalProfit += v.ExpectedAmount
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to get saving order", "Амжилттай"}, output))
}

func getRatio(term int64) float32 {
	var ratio float32

	if term == 6 {
		ratio = 0.015
	} else if term == 12 {
		ratio = 0.017
	} else if term == 24 {
		ratio = 0.019
	} else if term >= 36 {
		ratio = 0.02
	}

	return ratio
}

func sendInvoice(sadp saving.SavingOrderPayment) (payment.QPayInvoiceResponse, error) {

	var ttk payment.QPayToken
	ttk.Last()

	input := payment.QPayInvoiceInput{
		InvoiceCode:         sadp.InvoiceCode,
		SenderInvoiceNo:     sadp.SenderInvoiceNo,
		InvoiceReceiverCode: sadp.InvoiceReceiverCode,
		InvoiceDescription:  sadp.InvoiceDescription,
		SenderBranchCode:    sadp.SenderBranchCode,
		Amount:              sadp.Amount,
		CallbackURL:         sadp.CallbackURL,
	}

	jsonData, err := json.Marshal(input)

	if err != nil {
		return payment.QPayInvoiceResponse{}, err
	}

	url := "https://merchant.qpay.mn/v2/invoice"
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

	fmt.Println("amount", input.Amount)

	var res payment.QPayInvoiceResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return payment.QPayInvoiceResponse{}, err
	}

	if err := config.DB.Model(&saving.SavingOrderPayment{}).Where("id = ?", sadp.ID).Update("invoice_id", res.InvoiceID).Error; err != nil {
		return payment.QPayInvoiceResponse{}, err
	}

	return res, nil

}
