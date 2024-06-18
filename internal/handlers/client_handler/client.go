package client_handler

import (
	"fmt"
	"kcloudb1/internal/config"
	"kcloudb1/internal/middleware"
	"kcloudb1/internal/models/client"
	"kcloudb1/internal/utils"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SendOTP(phone, message string) error {

	url := fmt.Sprintf("https://api.messagepro.mn/send?to=%v&from=72887388&text=%v", phone, message)

	fmt.Println("url", url)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-api-key", "e15b92d6da557174aeb74b29f5243f77")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp)
	return err
}

func GenerateOTP(c *gin.Context) {
	var data struct {
		Phone string `json:"phone"`
	}

	var client, clientprev client.Client
	var resp utils.Response
	if err := c.ShouldBindJSON(&data); err != nil {
		resp = utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// otp generate and sent
	otp := fmt.Sprintf("%d", 1000+rand.Intn(8999))

	client.Phone = data.Phone
	client.OTP = otp
	client.OTPExpire = time.Now().Add(5 * time.Minute)

	if err := clientprev.GetByPhone(data.Phone); err != nil {
		fmt.Println("err", err)
		if err := client.Create(); err != nil {
			resp = utils.Error([]string{"Failed to create client phone and otp", "Алдаа гарлаа"}, err)
			c.JSON(http.StatusInternalServerError, resp)
			return
		}

	} else {
		client.ID = clientprev.ID
		if err := client.Update(); err != nil {
			resp = utils.Error([]string{"Failed to update client phone and otp", "Алдаа гарлаа"}, err)
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
	}

	err := SendOTP(data.Phone, fmt.Sprintf("Нэвтрэх%%20код:%s", otp))

	if err != nil {
		resp = utils.Error([]string{"Failed to send otp", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp = utils.Success([]string{"Success to send otp", "Амжилттай"}, nil)
	c.JSON(http.StatusOK, resp)
}

func VerifyOTP(c *gin.Context) {
	var data struct {
		Phone string `json:"phone"`
		Otp   string `json:"otp"`
	}

	var client client.Client
	var resp utils.Response
	if err := c.ShouldBindJSON(&data); err != nil {
		resp = utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	client.Phone = data.Phone

	if err := client.GetByPhone(data.Phone); err != nil {
		resp = utils.Error([]string{"Failed to get client by phone", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if client.OTP != data.Otp {
		resp = utils.Error([]string{"Invalid OTP", "Алдаа гарлаа"}, fmt.Errorf("Invalid OTP"))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if client.OTPExpire.Before(time.Now()) {
		resp = utils.Error([]string{"OTP expired", "Алдаа гарлаа"}, fmt.Errorf("OTP expired"))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	client.OTP = ""
	client.IsActive = 1

	if err := client.Update(); err != nil {
		resp = utils.Error([]string{"Failed to update client active status", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp = utils.Success([]string{"Success to verify otp", "Амжилттай"}, nil)
	c.JSON(http.StatusOK, resp)
}

func Register(c *gin.Context) {
	var data struct {
		Phone     string `json:"phone"`
		Password  string `json:"password"`
		Pin       string `json:"pin"`
		Firstname string `json:"first_name"`
		Lastname  string `json:"last_name"`
	}

	var client client.Client
	var resp utils.Response
	if err := c.ShouldBindJSON(&data); err != nil {
		resp = utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if err := client.GetByPhone(data.Phone); err != nil {
		resp = utils.Error([]string{"Failed to get client by phone", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if client.IsRegistered == 1 {
		resp = utils.Error([]string{"Phone already registered", "Алдаа гарлаа"}, nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	client.Password = data.Password
	client.Pin = data.Pin
	client.FirstName = data.Firstname
	client.LastName = data.Lastname
	client.IsRegistered = 1

	if err := client.Update(); err != nil {
		resp = utils.Error([]string{"Failed to update client", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp = utils.Success([]string{"Success to register", "Амжилттай"}, nil)
	c.JSON(http.StatusOK, resp)
}

func CheckToken(c *gin.Context) {
	var token = c.Query("token")

	if token == "" {
		c.JSON(http.StatusUnauthorized, utils.Error(
			[]string{"Unauthorized", "Нэвтрээгүй байна"},
			"Unauthorized",
		))
		c.Abort()
		return
	}

	claims, err := middleware.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.Error(
			[]string{"Unauthorized", "Нэвтрээгүй байна"},
			err.Error(),
		))
		c.Abort()
		return
	}

	fmt.Println("exp : ", claims["exp"].(float64))

	tokenExpire := time.Unix(int64(claims["exp"].(float64)), 0)
	if time.Now().After(tokenExpire) {
		c.JSON(http.StatusUnauthorized, utils.Error(
			[]string{"Unauthorized", "Нэвтрээгүй байна"},
			"Unauthorized",
		))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to check token", "Амжилттай"}, nil))

}

func LoginByPassword(c *gin.Context) {
	var data struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	var clientd client.Client
	var clientOutput client.ClientOutput
	var resp utils.Response

	if err := c.ShouldBindJSON(&data); err != nil {
		resp = utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	clientd.Phone = data.Phone
	clientd.Password = data.Password

	if err := clientd.GetByPhone(data.Phone); err != nil {
		resp = utils.Error([]string{"Failed to get client by phone", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if clientd.IsRegistered == 0 {
		resp = utils.Error([]string{"Phone not registered", "Алдаа гарлаа"}, fmt.Errorf("Phone not registered"))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if clientd.Password != data.Password {
		resp = utils.Error([]string{"Invalid password", "Алдаа гарлаа"}, fmt.Errorf("Invalid password"))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	token := uuid.NewString()

	result := config.RS.Set(token, fmt.Sprintf("%d", clientd.ID), 0)

	if result.Err() != nil {
		resp = utils.Error([]string{"Failed to create token", "Алдаа гарлаа"}, result.Err())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	tokenJWT, err := middleware.CreateToken(clientd.ID, token)

	if err != nil {
		resp = utils.Error([]string{"Failed to create token", "Алдаа гарлаа"}, err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	clientOutput.ID = clientd.ID
	clientOutput.FirstName = clientd.FirstName
	clientOutput.LastName = clientd.LastName
	clientOutput.Phone = clientd.Phone
	clientOutput.IsActive = clientd.IsActive
	clientOutput.IsRegistered = clientd.IsRegistered
	clientOutput.CreatedAt = clientd.CreatedAt
	clientOutput.Token = tokenJWT

	resp = utils.Success([]string{"Success to login", "Амжилттай"}, clientOutput)
	c.JSON(http.StatusOK, resp)

}

func ForgotPassword(c *gin.Context) {
	var input struct {
		Phone string `json:"phone"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err))
		return
	}

	otp := fmt.Sprintf("%d", 1000+rand.Intn(8999))

	client := client.Client{Phone: input.Phone}

	if err := client.GetByPhone(input.Phone); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to get client by phone", "Алдаа гарлаа"}, err))
		return
	}

	if client.IsRegistered == 0 {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Phone not registered", "Алдаа гарлаа"}, fmt.Errorf("Phone not registered")))
		return
	}

	client.OTP = otp
	client.OTPExpire = time.Now().Add(5 * time.Minute)

	if err := client.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to save otp", "Алдаа гарлаа"}, err))
		return
	}

	if err := client.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to update otp", "Алдаа гарлаа"}, err))
		return
	}

	if err := SendOTP(input.Phone, otp); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to send otp", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to send otp", "Амжилттай"}, nil))
}

func VerifyOTPChangePassword(c *gin.Context) {
	var input struct {
		Phone    string `json:"phone"`
		Otp      string `json:"otp"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err))
		return
	}

	client := client.Client{Phone: input.Phone}
	if err := client.GetByPhone(input.Phone); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to get client by phone", "Алдаа гарлаа"}, err))
		return
	}

	if client.OTP != input.Otp {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Invalid OTP", "Алдаа гарлаа"}, fmt.Errorf("Invalid OTP")))
		return
	}

	if client.OTPExpire.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"OTP expired", "Алдаа гарлаа"}, fmt.Errorf("OTP expired")))
		return
	}

	client.Password = input.Password
	if err := client.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to update password", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to change password", "Амжилттай"}, nil))

}

func ChangePassword(c *gin.Context) {

	var input struct {
		Phone       string `json:"phone"`
		Password    string `json:"password"`
		NewPassword string `json:"new_password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Failed to bind json", "Алдаа гарлаа"}, err))
		return
	}

	client := client.Client{Phone: input.Phone}
	if err := client.GetByPhone(input.Phone); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to get client by phone", "Алдаа гарлаа"}, err))
		return
	}

	if input.Password != client.Password {
		c.JSON(http.StatusBadRequest, utils.Error([]string{"Invalid password", "Алдаа гарлаа"}, fmt.Errorf("Invalid password")))
		return
	}

	client.Password = input.NewPassword
	if err := client.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error([]string{"Failed to update password", "Алдаа гарлаа"}, err))
		return
	}

	c.JSON(http.StatusOK, utils.Success([]string{"Success to change password", "Амжилттай"}, nil))
}
