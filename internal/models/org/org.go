package org

import "time"

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
