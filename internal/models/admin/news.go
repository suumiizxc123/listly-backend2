package admin

import (
	"kcloudb1/internal/config"
	"time"
)

type News struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	Image     string    `json:"image"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *News) TableName() string {
	return "one_news"
}

func (c *News) Create() error {
	return config.DB.Create(c).Error
}

func (c *News) Get() error {
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *News) Update() error {
	return config.DB.Updates(c).Error
}

func (c *News) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *News) GetAll() ([]News, error) {
	var news []News
	if err := config.DB.Order("created_at desc").Find(&news).Error; err != nil {
		return nil, err
	}
	return news, nil
}
