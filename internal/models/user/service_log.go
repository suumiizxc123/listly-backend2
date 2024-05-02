package user

import (
	"kcloudb1/internal/config"
	"time"
)

type ServiceLog struct {
	ID        int64     `json:"ID" gorm:"primary_key"`
	KaraokeID int64     `json:"karaoke_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *ServiceLog) TableName() string {
	return "service_log"
}

func (c *ServiceLog) Create() error {
	return config.DB.Create(c).Error
}

func (c *ServiceLog) GetList() ([]ServiceLog, error) {
	var serviceLogs []ServiceLog

	if err := config.DB.Find(&serviceLogs).Error; err != nil {
		return nil, err
	}

	return serviceLogs, nil
}

func (c *ServiceLog) Update() error {
	return config.DB.Updates(c).Error
}

func (c *ServiceLog) Delete() error {
	return config.DB.Delete(c).Error
}
