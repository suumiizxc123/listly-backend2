package moscap

import (
	"kcloudb1/internal/config"
	"time"
)

type MosCapUserLog struct {
	ID           int64     `json:"ID" gorm:"primary_key"`
	MosCapUserID int64     `json:"moscap_user_id"`
	Title        string    `json:"title"`
	ActionID     int64     `json:"action_id"`
	CreatedAt    time.Time `json:"created_at"`
}

func (c *MosCapUserLog) TableName() string {
	return "moscap_user_log"
}

func (c *MosCapUserLog) Create() error {
	return config.DB.Create(c).Error
}

func (c *MosCapUserLog) Update() error {
	return config.DB.Updates(c).Error
}

func (c *MosCapUserLog) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *MosCapUserLog) Get() error {
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *MosCapUserLog) GetList() ([]MosCapUserLog, error) {
	var orgs []MosCapUserLog

	if err := config.DB.Find(&orgs).Error; err != nil {
		return nil, err
	}

	return orgs, nil
}
