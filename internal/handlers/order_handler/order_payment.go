package order_handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"kcloudb1/internal/config"
	"kcloudb1/internal/models/order"
	"kcloudb1/internal/models/payment"
	"kcloudb1/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckPaymentCallBack(c *gin.Context) {
	newUID := c.Param("newuid")

	if newUID == "" {
		c.String(http.StatusBadRequest, "New UID is empty")
		return
	}

	var ordp order.OrderPayment

	if err := config.DB.Where("sender_invoice_no = ?", newUID).First(&ordp).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.String(http.StatusNotFound, "Order payment not found")
			return
		}

		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err := config.DB.Model(&order.Order{}).Where("id = ?", ordp.OrderID).Update("status", "success").Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	var ord order.Order

	if err := config.DB.Where("id = ?", ordp.OrderID).First(&ord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.String(http.StatusNotFound, "Order not found")
			return
		}

		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	var bal order.Balance

	if err := bal.GetByClientAndMetalID(ord.ClientID, 1); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err := config.DB.Model(&order.Balance{}).Where("id = ?", bal.ID).Update("quantity", bal.Quantity+ord.Quantity).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, newUID)

}

func CheckPayment(c *gin.Context) {

	newUID := c.Param("newuid")

	if newUID == "" {
		resp := utils.Error([]string{"New UID is empty", "Алдаа гарлаа"}, fmt.Errorf("new uid is empty"))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	var ordp order.OrderPayment

	if err := config.DB.Where("sender_invoice_no = ?", newUID).First(&ordp).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp := utils.Error([]string{"Record not found", "Алдаа гарлаа"}, fmt.Errorf("record not found"))
			c.JSON(http.StatusNotFound, resp)
			return
		}
	}

	chk, err := checkQPaymentCallBack(ordp.InvoiceID, ordp.Amount)

	if err != nil {
		resp := utils.Error([]string{"Failed to check payment", "Алдаа гарлаа"}, err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := utils.Success([]string{"Success to check payment", "Амжилттай"}, chk)
	c.JSON(http.StatusOK, resp)
}

func checkQPaymentCallBack(invoiceID string, amount float32) (payment.QPayCheck, error) {
	var ttk payment.QPayToken
	var qpaychk payment.QPayCheck
	ttk.Last()

	var input struct {
		ObjectType string `json:"object_type"`
		ObjectID   string `json:"object_id"`
	}

	input.ObjectID = invoiceID

	input.ObjectType = "INVOICE"

	jsonData, err := json.Marshal(input)

	if err != nil {
		return payment.QPayCheck{}, err
	}

	fmt.Println("jsonData:", string(jsonData))

	url := "https://merchant.qpay.mn/v2/payment/check"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return payment.QPayCheck{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+ttk.AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return payment.QPayCheck{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return payment.QPayCheck{}, err
	}

	err = json.Unmarshal(body, &qpaychk)
	if err != nil {
		return payment.QPayCheck{}, err
	}

	fmt.Println("qpaychk:", qpaychk)

	if qpaychk.Count == 0 {
		return payment.QPayCheck{}, errors.New("payment not found")
	}

	if qpaychk.PaidAmount == 0 {
		return payment.QPayCheck{}, errors.New("payment not found")
	}

	if qpaychk.PaidAmount != amount {
		return payment.QPayCheck{}, errors.New("payment not completed")
	}

	return qpaychk, nil

}
