package order_handler

import (
	"errors"
	"kcloudb1/internal/config"
	"kcloudb1/internal/models/order"
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
			c.String(http.StatusNotFound, "Record not found")
			return
		}

		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err := config.DB.Model(&order.Order{}).Where("id = ?", ordp.OrderID).Update("status", "success").Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, newUID)
}
