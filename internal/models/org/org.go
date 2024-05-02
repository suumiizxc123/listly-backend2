package org

import (
	"kcloudb1/internal/config"
	"time"
)

type Org struct {
	ID          int64     `json:"ID" gorm:"primary_key"`
	KaraokeName string    `json:"karaoke_name"`
	Address     string    `json:"address"`
	Picture     string    `json:"picture"`
	Latitude    string    `json:"latitude"`
	Longitude   string    `json:"longitude"`
	PhoneNumber string    `json:"phone_number"`
	IsActive    int64     `json:"is_active"`
	Rating      float32   `json:"rating"`
	CreatedAt   time.Time `json:"created_at"`
	ExpireDate  time.Time `json:"expire_date"`
	AuthToken   string    `json:"auth_token"`
}

func (c *Org) TableName() string {
	return "org"
}

func (c *Org) Create() error {
	return config.DB.Create(c).Error
}

func (c *Org) Update() error {
	return config.DB.Updates(c).Error
}

func (c *Org) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *Org) Get() error {
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *Org) GetList() ([]Org, error) {
	var orgs []Org

	if err := config.DB.Find(&orgs).Error; err != nil {
		return nil, err
	}

	return orgs, nil
}
