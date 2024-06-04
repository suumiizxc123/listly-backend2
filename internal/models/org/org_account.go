package org

import "kcloudb1/internal/config"

type OrgAccount struct {
	ID        int64   `json:"ID" gorm:"primary_key"`
	KaraokeID int64   `json:"karaoke_id"`
	Balance   float32 `json:"balance"`
}

func (c *OrgAccount) TableName() string {
	return "org_account"
}

func (c *OrgAccount) Create() error {
	return config.DB.Create(c).Error
}

func (c *OrgAccount) Update() error {
	return config.DB.Updates(c).Error
}

func (c *OrgAccount) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *OrgAccount) Get() error {
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *OrgAccount) GetList() ([]OrgAccount, error) {
	var orgs []OrgAccount

	if err := config.DB.Find(&orgs).Error; err != nil {
		return nil, err
	}

	return orgs, nil
}
