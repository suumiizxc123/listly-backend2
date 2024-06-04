package client_handler

import (
	"fmt"
	"kcloudb1/internal/models/client"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SendOTP(phone, message string) error {

	url := fmt.Sprintf("https://api.messagepro.mn/send?to=%v&from=72022001&text=%v", phone, message)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("User-Agent", "insomnia/9.2.0")
	req.Header.Add("x-api-key", "e6d4f2ca52e24896f3f238be8df4fbc8")

	_, err := http.DefaultClient.Do(req)
	return err
}

func GenerateOTP(c *gin.Context) {
	var data struct {
		Phone string `json:"phone"`
	}

	var client client.Client

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// otp generate and sent
	otp := fmt.Sprintf("%d", 1000+rand.Intn(8999))

	client.Phone = data.Phone
	client.OTP = otp
	client.IsActive = 0
	client.OTPExpire = time.Now().Add(5 * time.Minute)

	if err := client.GetByPhone(data.Phone); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone already registered", "message": "Phone already registered"})
		return
	}

	if err := client.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to create client phone and otp"})
		return
	}

	err := SendOTP(data.Phone, otp)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to send otp"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent"})
}

func VerifyOTP(c *gin.Context) {
	var data struct {
		Phone string `json:"phone"`
		Otp   string `json:"otp"`
	}

	var client client.Client

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client.Phone = data.Phone

	if err := client.GetByPhone(data.Phone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if client.OTP != data.Otp {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OTP", "message": "Invalid OTP"})
		return
	}

	if client.OTPExpire.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "OTP expired", "message": "OTP expired"})
		return
	}

	if client.IsActive == 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "OTP already verified", "message": "OTP already verified"})
		return
	}

	client.OTP = uuid.NewString()
	client.IsActive = 1

	if err := client.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP verified", "token": client.OTP})
}

func Register(c *gin.Context) {
	var data struct {
		Token     string `json:"token"`
		Password  string `json:"password"`
		Pin       string `json:"pin"`
		Firstname string `json:"first_name"`
		Lastname  string `json:"last_name"`
	}

	var client client.Client

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client.OTP = data.Token

	if err := client.GetByOTP(data.Token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if client.IsRegistered == 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone already registered", "message": "Phone already registered"})
		return
	}

	client.Password = data.Password
	client.Pin = data.Pin
	client.FirstName = data.Firstname
	client.LastName = data.Lastname
	client.IsRegistered = 1

	if err := client.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Register successful"})

}

func LoginByPassword(c *gin.Context) {
	var data struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	var client client.Client

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client.Phone = data.Phone
	client.Password = data.Password

	if err := client.GetByPhone(data.Phone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if client.IsRegistered == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone not registered", "message": "Phone not registered"})
		return
	}

	if client.Password != data.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password", "message": "Invalid password"})
		return
	}

	token := uuid.NewString()

	client.Token = token

	if err := client.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})

}
