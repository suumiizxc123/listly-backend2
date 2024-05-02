package org

import (
	"kcloudb1/internal/config"
	"kcloudb1/internal/models/user"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"
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

type OrgAndUserInput struct {
	KaraokeName string `json:"karaoke_name"`
	Address     string `json:"address"`
	Picture     string `json:"picture"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	PhoneNumber string `json:"phone_number"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}

func (c *OrgAndUserInput) Create() error {

	tx := config.DB.Begin()
	org := Org{
		KaraokeName: c.KaraokeName,
		Address:     c.Address,
		Picture:     c.Picture,
		Latitude:    c.Latitude,
		Longitude:   c.Longitude,
		PhoneNumber: c.PhoneNumber,
		Rating:      0.0,
		IsActive:    1,
		CreatedAt:   time.Now(),
		ExpireDate:  time.Now().AddDate(1, 0, 0),
	}

	if err := tx.Create(&org).Error; err != nil {
		tx.Rollback()
		return err
	}

	user := user.User{
		UID:       uuid.New().String(),
		RoleID:    1,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		KaraokeID: org.ID,
		Email:     c.Email,
		Phone:     c.Phone,
		Password:  c.Password,
		IsActive:  1,
		CreatedAt: time.Now(),
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error

}

type OrgExtend struct {
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
	User        user.User `json:"user" gorm:"foreignKey:KaraokeID; references:ID"`
}

func (c *OrgExtend) TableName() string {
	return "org"
}

func (c *OrgExtend) GetList(offset int, limit int, order string, sort string) ([]OrgExtend, error) {
	var orgs []OrgExtend

	if err := config.DB.
		Offset(offset).Limit(limit).Order(sort + " " + order).Preload(clause.Associations).Find(&orgs).Error; err != nil {

		return nil, err
	}

	return orgs, nil
}
