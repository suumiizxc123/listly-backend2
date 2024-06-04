package org

import (
	"kcloudb1/internal/config"
	"time"
)

type OrgAccountTxnLog struct {
	ID         int64     `json:"ID" gorm:"primary_key"`
	TxnID      int64     `json:"txn_id"`
	AccountID  int64     `json:"account_id"`
	BegBal     float32   `json:"beg_bal"`
	Amount     float32   `json:"amount"`
	EndBal     float32   `json:"end_bal"`
	SysUserID  int64     `json:"sys_user_id"`
	ChargeType int64     `json:"charge_type"`
	CreatedAt  time.Time `json:"created_at"`
}

func (c *OrgAccountTxnLog) TableName() string {
	return "org_account_txn_log"
}

func (c *OrgAccountTxnLog) Create() error {
	return config.DB.Create(c).Error
}

func (c *OrgAccountTxnLog) Update() error {
	return config.DB.Updates(c).Error
}

func (c *OrgAccountTxnLog) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *OrgAccountTxnLog) Get() error {
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *OrgAccountTxnLog) GetList() ([]OrgAccountTxnLog, error) {
	var orgs []OrgAccountTxnLog

	if err := config.DB.Find(&orgs).Error; err != nil {
		return nil, err
	}

	return orgs, nil
}
