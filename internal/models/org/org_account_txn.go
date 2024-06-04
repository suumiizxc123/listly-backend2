package org

import (
	"kcloudb1/internal/config"
	"time"
)

type OrgAccountTxn struct {
	ID         int64     `json:"ID" gorm:"primary_key"`
	AccountID  int64     `json:"account_id"`
	Amount     float32   `json:"amount"`
	SysUserID  int64     `json:"sys_user_id"`
	ChargeType int64     `json:"charge_type"`
	CreatedAt  time.Time `json:"created_at"`
}

func (c *OrgAccountTxn) TableName() string {
	return "org_account_txn"
}

func (c *OrgAccountTxn) Create() error {
	return config.DB.Create(c).Error
}

func (c *OrgAccountTxn) Update() error {
	return config.DB.Updates(c).Error
}

func (c *OrgAccountTxn) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *OrgAccountTxn) Get() error {
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *OrgAccountTxn) GetList() ([]OrgAccountTxn, error) {
	var orgs []OrgAccountTxn

	if err := config.DB.Find(&orgs).Error; err != nil {
		return nil, err
	}

	return orgs, nil
}
