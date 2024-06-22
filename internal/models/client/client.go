package client

import (
	"kcloudb1/internal/config"
	"time"
)

type Client struct {
	ID           int64     `json:"ID" gorm:"primary_key"`
	Phone        string    `json:"phone"`
	Password     string    `json:"password"`
	Pin          string    `json:"pin"`
	IsActive     int64     `json:"is_active"`
	IsRegistered int64     `json:"is_registered"`
	OTP          string    `json:"otp"`
	OTPExpire    time.Time `json:"otp_expire"`
	Token        string    `json:"token"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	CreatedAt    time.Time `json:"created_at"`
}

func (c *Client) TableName() string {
	return "one_client"
}

func (c *Client) Create() error {
	return config.DB.Create(c).Error
}

func (c *Client) Update() error {
	return config.DB.Updates(c).Error
}

func (c *Client) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *Client) Get() error {
	return config.DB.First(c, c.ID).Error
}

func (c *Client) GetAll() ([]Client, error) {
	var clients []Client
	err := config.DB.Find(&clients).Error
	return clients, err
}

func (c *Client) GetByPhone(phone string) error {
	return config.DB.Where("phone = ?", phone).First(c).Error
}

func (c *Client) Save(pin string) error {
	return config.DB.Save(c).Error
}

func (c *Client) GetByOTP(otp string) error {
	return config.DB.Where("otp = ?", otp).First(c).Error
}

type ClientOutput struct {
	ID           int64     `json:"ID"`
	Phone        string    `json:"phone"`
	IsActive     int64     `json:"is_active"`
	IsRegistered int64     `json:"is_registered"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	CreatedAt    time.Time `json:"created_at"`
	Token        string    `json:"token"`
}
