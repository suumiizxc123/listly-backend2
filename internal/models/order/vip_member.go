package order

import (
	"kcloudb1/internal/config"
	"time"
)

type VipMember struct {
	ID        int64     `json:"id"`
	ClientID  int64     `json:"client_id"`
	Amount    float32   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *VipMember) TableName() string {
	return "one_vip_member"
}

func (c *VipMember) Create() error {
	return config.DB.Create(c).Error
}

func (c *VipMember) Get() error {
	return config.DB.First(c, c.ID).Error
}

func (c *VipMember) Update() error {
	return config.DB.Updates(c).Error
}

func (c *VipMember) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *VipMember) GetByClientID() error {
	return config.DB.Where("client_id = ?", c.ClientID).First(c).Error
}

func (c *VipMember) GetList() ([]VipMember, error) {
	var vip []VipMember
	return vip, config.DB.Find(&vip).Error
}

func (c *VipMember) GetListByClientID() ([]VipMember, error) {
	var vip []VipMember
	return vip, config.DB.Where("client_id = ?", c.ClientID).Find(&vip).Error
}
