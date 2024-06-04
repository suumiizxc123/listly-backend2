package org

import (
	"kcloudb1/internal/config"
	"time"
)

type OrgUser struct {
	ID        int64     `json:"ID" gorm:"primary_key"`
	UID       string    `json:"uid"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	KaraokeID int64     `json:"karaoke_id"`
	RoleID    int64     `json:"role_id"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	IsActive  int64     `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	Token     string    `gorm:"-" json:"token, omitempty"`
}

func (c *OrgUser) TableName() string {
	return "org_user"
}

func (c *OrgUser) Create() error {
	return config.DB.Create(c).Error
}

func (c *OrgUser) Update() error {
	return config.DB.Updates(c).Error
}

func (c *OrgUser) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *OrgUser) Get() error {
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *OrgUser) GetList() ([]OrgUser, error) {
	var orgs []OrgUser

	if err := config.DB.Find(&orgs).Error; err != nil {
		return nil, err
	}

	return orgs, nil
}
