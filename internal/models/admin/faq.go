package admin

import (
	"kcloudb1/internal/config"
	"time"
)

type FAQ struct {
	ID        int64     `json:"ID"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
	IsActive  int64     `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *FAQ) TableName() string {
	return "admin_faq"
}

func (c *FAQ) Create() error {
	return config.DB.Create(c).Error
}

func (c *FAQ) Get() error {
	return config.DB.First(c, c.ID).Error
}

func (c *FAQ) Update() error {
	return config.DB.Updates(c).Error
}

func (c *FAQ) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *FAQ) GetList() ([]FAQ, error) {
	var list []FAQ
	err := config.DB.Order("created_at desc").Find(&list).Error
	return list, err
}
